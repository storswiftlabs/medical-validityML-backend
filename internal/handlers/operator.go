package handlers

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"medical-zkml-backend/internal/db"
	m "medical-zkml-backend/internal/module"
	"medical-zkml-backend/internal/plugin/exec/python"
	"medical-zkml-backend/internal/plugin/prediction_module_interface"
	"medical-zkml-backend/pkg/config"
	"medical-zkml-backend/pkg/utils"
	"os"
	"strings"
	"time"
)

type Operator struct {
	Modules []m.Module
}

func (o *Operator) InModuleList(module string) bool {
	modules := o.Modules
	for _, v := range modules {
		if v.Name == module {
			return true
		}
	}

	return false
}

// FIXME: Normalization without using Python
func (o *Operator) normalization(disease string, params []string) ([]string, error) {
	data, err := python.NormalizedData(disease, params)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Normalized data execution failed: %s", err.Error()))
	}

	return data, nil
}

func (o *Operator) quantization(params []string) []int {
	inputs := utils.Strings2Floats(params)
	list := make([]int, 0, len(inputs))

	for _, input := range inputs {
		list = append(list, int(input*65535))
	}

	return list
}

func (o *Operator) DiseasePrediction(premise *m.PredictionPremise) {

	zap.L().Info(fmt.Sprintf("%s started %s disease prediction", premise.User, premise.Name))

	// Write this prediction to the database in the start state
	zap.L().Info(fmt.Sprintf("%s started information record", premise.User))
	recordPrediction := &m.RecordPrediction{
		User:          premise.User,
		Status:        m.Start,
		Disease:       premise.Name,
		Module:        premise.Module,
		Inputs:        utils.CoverInput(premise, "key"),
		InputsMapping: utils.CoverInput(premise, "value"),
		StartTime:     time.Now().Unix(),
		EndTime:       time.Now().Add(3 * time.Minute).Unix(),
		Output:        -1,
	}
	if err := db.RecordPredict(recordPrediction); err != nil {
		zap.L().Error(fmt.Sprintf("%s information record failed", premise.User), zap.Error(err))
	}
	zap.L().Info(fmt.Sprintf("%s information record successful", premise.User))

	params := make([]string, len(premise.Inputs))
	for _, input := range premise.Inputs {
		params[input.Index] = input.SelectValue
	}

	zap.L().Info(fmt.Sprintf("%s disease information %v", premise.User, params))

	// Data normalization processing
	zap.L().Info(fmt.Sprintf("%s started data normalization", premise.User))
	normalized, err := o.normalization(premise.Name, params)
	if err != nil {
		zap.L().Error(fmt.Sprintf("%s data normalization failed", premise.User), zap.Error(err))
		recordPrediction.Status = m.NormalizationFailed
		recordPrediction.Message = fmt.Sprintf("%s: %s", "Data quantization failure", err.Error())
		if err := db.RecordPredict(recordPrediction); err != nil {
			zap.L().Error(fmt.Sprintf("Data storage failure: %s", err.Error()))
			return
		}
		return
	}
	zap.L().Info(fmt.Sprintf("%s data normalization successful", premise.User), zap.Any("data", normalized))
	recordPrediction.Status = m.NormalizationSuccess
	recordPrediction.Normalized = strings.Join(normalized, ", ")
	if err := db.RecordPredict(recordPrediction); err != nil {
		zap.L().Error(fmt.Sprintf("Data storage failure: %s", err.Error()))
		return
	}

	// Data quantization processing
	zap.L().Info(fmt.Sprintf("%s started data quantization", premise.User))
	quantized := o.quantization(normalized)
	zap.L().Info(fmt.Sprintf("%s data quantization successful", premise.User), zap.Any("data", quantized))

	recordPrediction.Quantized = utils.Ints2String(quantized)
	if err := db.RecordPredict(recordPrediction); err != nil {
		zap.L().Error(fmt.Sprintf("Data storage failure: %s", err.Error()))
		return
	}

	// disease prediction
	prediction := prediction_module_interface.GetPredictionModule()
	prediction.DiseasePrediction(recordPrediction, premise.Name, premise.Module, quantized)

	return
}

func (o *Operator) GetProposeByName(disease string) (string, error) {
	return db.GetRecommendation(disease)
}

func (o *Operator) GetPrediction(user string, page, pageSize int) ([]m.PredictedResult, int64, error) {
	count, err := db.GetPredictedCount(user)
	if err != nil {
		return nil, 0, err
	}
	list, err := db.GetPredictedList(user, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

func (o *Operator) VerifyInformation(req *m.VerifyReq) (*m.VerifyResultNoModel, error) {
	result, err := db.VerifyInformation(req)
	if err != nil {
		return nil, err
	}
	result.ContractAddress = config.GetConfig().Sub("contract").Sub(fmt.Sprintf("%s+%s", strings.ReplaceAll(result.Module, " ", "_"), strings.ReplaceAll(result.Disease, " ", "_"))).Get("address").(string)
	result.ContractFunction = "verify"
	data, err := os.ReadFile(fmt.Sprintf("internal/plugin/abi/%s/%s.json", result.Module, strings.ReplaceAll(result.Disease, " ", "_")))
	if err != nil {
		return nil, err
	}
	result.ABI = string(data)
	return result, nil
}

func (o *Operator) GetArticle(diseases []string) ([]m.ArticleResult, error) {
	for i := 0; i < len(diseases); i++ {
		diseases[i] = strings.ReplaceAll(diseases[i], " ", "_")
	}
	return db.GetArticle(diseases)
}

func (o *Operator) DeletePredictedResults(user string, ids []int) error {
	return db.DeletePredictedForUser(user, ids)
}

func (o *Operator) GetPredicted(user string, id uint) (m.RecordPrediction, error) {
	return db.GetPredicted(user, id)
}

func (o *Operator) Reload(disease *m.DiseaseInfo, kv []m.KeyValue) *m.DiseaseInfoResult {
	inputs := make([]m.InputResult, len(disease.Inputs))
	for index, input := range disease.Inputs {
		inputs[index].Name = input.Name
		inputs[index].Description = input.Description
		inputs[index].Index = input.Index
		inputs[index].InputMax = input.InputMax
		inputs[index].InputMin = input.InputMin
		inputs[index].InputDecimalLength = input.InputDecimalLength
		inputs[index].Warn = input.Warn
		inputs[index].InputMethod = input.InputMethod
		inputs[index].Select = input.Select
		for _, v := range kv {
			if input.Name == v.Key.(string) {
				inputs[index].Default = &v
				break
			}
		}
	}

	return &m.DiseaseInfoResult{
		Name:        disease.Name,
		Description: disease.Description,
		InputResult: inputs,
	}
}

func (o *Operator) PredictingOutcomesFuzzyQuery(user, key string, page, pageSize int) ([]m.PredictedResult, int64, error) {
	count, err := db.GetPredictedFuzzyQueryCount(user, key)
	if err != nil {
		return nil, 0, err
	}

	result, err := db.PredictingOutcomesFuzzyQuery(user, key, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return result, count, err
}
