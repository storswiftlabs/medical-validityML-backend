package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"medical-zkml-backend/internal/db"
	"medical-zkml-backend/internal/module"
	"medical-zkml-backend/pkg/utils"
	"sort"
	"strconv"
	"strings"
)

type Disease struct {
	diseaseList []module.Disease
	diseaseInfo module.DiseaseList
}

func NewDisease(diseaseList []module.Disease, diseaseInfo module.DiseaseList) Disease {
	return Disease{
		diseaseList: diseaseList,
		diseaseInfo: diseaseInfo,
	}
}

func (d *Disease) GetDiseaseNameList() []string {
	var names []string
	list := d.diseaseList
	for _, disease := range list {
		names = append(names, disease.Name)
	}
	return names
}

func (d *Disease) elemInArray(elem string) bool {
	list := d.diseaseList
	for _, disease := range list {
		if disease.Name == elem {
			return true
		}
	}
	return false
}

func (d *Disease) CountInputsForName(name string) int {
	return len(d.diseaseInfo[name].Inputs)
}

func (d *Disease) GetDescriptionByName(disease string) string {
	list := d.diseaseList
	for _, v := range list {
		if disease == v.Name {
			return v.Description
		}
	}
	return ""
}

func (d *Disease) GetOutputValue(disease, key string) (string, error) {
	output := d.diseaseInfo[disease].Output.Result
	for _, kv := range output {
		if kv.Key == key {
			return kv.Value.(string), nil
		}
	}
	zap.L().Error(fmt.Sprintf("%s disease not found the key %s", disease, key))
	return "", errors.New("disease mapping search failed")
}

func (d *Disease) ParameterCheck(info *module.PredictionPremise) error {
	inputs := d.diseaseInfo[info.Name].Inputs

	for index, input := range inputs {
		if input.InputMethod == "input" {
			warn := fmt.Errorf("%s input error. you entered %s, %s", input.Name, info.Inputs[index].SelectValue, input.Warn)
			valueFloat, err := strconv.ParseFloat(info.Inputs[index].SelectValue, 64)
			if err != nil {
				return warn
			}

			var decimalLen int
			if strings.Contains(info.Inputs[index].SelectValue, ".") {
				decimalLen = len(info.Inputs[index].SelectValue) - strings.Index(info.Inputs[index].SelectValue, ".") - 1
			}
			if decimalLen > input.InputDecimalLength {
				return warn
			}
			if !(valueFloat >= input.InputMin && valueFloat <= input.InputMax) {
				return warn
			}
			InputPreprocessing(&input)
			continue
		}

		sort.Slice(input.Select, func(i, j int) bool {
			return input.Select[i].Value.(int) < input.Select[j].Value.(int)
		})

		choose, err := strconv.Atoi(info.Inputs[index].SelectValue)
		if err != nil {
			return err
		}

		if !(input.Select[0].Value.(int) <= choose && input.Select[len(input.Select)-1].Value.(int) >= choose) {
			return errors.New("the selected parameter is out of range")
		}
	}
	return nil
}

func InputPreprocessing(input *module.Input) {
	switch input.Name {
	case "Specific gravity":
		SpecificGravityPreprocessing(input)
	}
}

func SpecificGravityPreprocessing(input *module.Input) {
	value := utils.Interface2float64(input.Select[0].Value)
	if value >= 1.000 && value <= 1.007 {
		input.Select[0].Key = 1.005
		input.Select[0].Value = 1.005
		return
	} else if value >= 1.008 && value <= 1.012 {
		input.Select[0].Key = 1.010
		input.Select[0].Value = 1.010
		return
	} else if value >= 1.013 && value <= 1.017 {
		input.Select[0].Key = 1.015
		input.Select[0].Value = 1.015
		return
	} else if value >= 1.0018 && value <= 1.023 {
		input.Select[0].Key = 1.020
		input.Select[0].Value = 1.020
		return
	} else if value >= 1.024 {
		input.Select[0].Key = 1.025
		input.Select[0].Value = 1.025
		return
	}
}

func (d *Disease) GetDiseaseFromPosition(info *module.InitialDiagnosisInfo) []module.Disease {
	//result := d.GetDiseaseFromCondition(info.MedicalCondition)
	var result []module.Disease
	for _, disease := range d.diseaseList {
		for _, position := range disease.Positions {
			if find(info.SymptomList, position) {
				if !checkDiseaseInDiseases(result, disease.Name) {
					result = append(result, disease)
					break
				}
			}
		}
	}
	return result
}

func checkDiseaseInDiseases(list []module.Disease, name string) bool {
	for _, disease := range list {
		if disease.Name == name {
			return true
		}
	}
	return false
}

func find(data []module.SymptomList, key string) bool {
	for _, d := range data {

		if d.Position == key {
			return true
		}
	}
	return false
}

func (d *Disease) GetDiseaseFromCondition(symptoms module.MedicalCondition) []module.Disease {
	var result []module.Disease
	for _, disease := range d.diseaseList {
		info := d.diseaseInfo[disease.Name]
		for _, input := range info.Inputs {
			if (input.Name == "Diabetes" && symptoms.Diabetes == "yes") ||
				(input.Name == "Obese" && symptoms.Obese == "yes") ||
				(input.Name == "Hypertension" && symptoms.Hypertension == "yes") ||
				(input.Name == "Smoking" && symptoms.Smoking == "yes") ||
				(input.Name == "Anemia" && symptoms.Anemia == "yes") ||
				(input.Name == "Cholesterol" && symptoms.Cholesterol == "yes") {
				result = append(result, disease)
				break
			}
		}
	}
	return result
}

func (d *Disease) CacheInitialDiagnosisInfo(diseaseRequest *module.InitialDiagnosisInfo) (uint, error) {
	bytes, err := json.Marshal(diseaseRequest)
	if err != nil {
		return 0, err
	}
	return db.CacheInitialDiagnosisInfo(string(bytes))
}

func (d *Disease) GetInitialDiagnosisInfo(id uint) (*module.InitialDiagnosisInfo, error) {
	var cache module.InitialDiagnosisInfo
	bytes, err := db.GetInitialDiagnosisInfo(id)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(bytes), &cache); err != nil {
		return nil, err
	}

	return &cache, nil
}

func (d *Disease) CopyDisease2Result(name string, cache *module.InitialDiagnosisInfo) *module.DiseaseInfoResult {
	disease := d.diseaseInfo[name]
	inputs := make([]module.InputResult, len(disease.Inputs))
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
		switch inputs[index].Name {
		case "Age":
			inputs[index].Default = &module.KeyValue{Key: inputs[index].Name, Value: cache.Age}
		case "Age range":
			inputs[index].Default = FindAge(inputs[index].Select, cache.Age)
		case "Sex":
			inputs[index].Default = Find(inputs[index].Select, cache.Gender)
		case "Diabetes":
			inputs[index].Default = Find(inputs[index].Select, cache.MedicalCondition.Diabetes)
		case "Obese":
			inputs[index].Default = Find(inputs[index].Select, cache.MedicalCondition.Obese)
		case "Hypertension":
			inputs[index].Default = Find(inputs[index].Select, cache.MedicalCondition.Hypertension)
		case "Smoking":
			inputs[index].Default = Find(inputs[index].Select, cache.MedicalCondition.Smoking)
		case "Anemia":
			inputs[index].Default = Find(inputs[index].Select, cache.MedicalCondition.Anemia)
		case "Cholesterol":
			inputs[index].Default = Find(inputs[index].Select, cache.MedicalCondition.Cholesterol)
		}
	}
	return &module.DiseaseInfoResult{
		Name:        disease.Name,
		Description: disease.Description,
		InputResult: inputs,
	}
}

func Find(data []module.KeyValue, key string) *module.KeyValue {
	for _, kv := range data {
		if strings.ToLower(kv.Key.(string)) == key {
			return &kv
		}
	}
	return nil
}

func FindAge(data []module.KeyValue, key int) *module.KeyValue {
	if key >= 90 {

		return &data[len(data)-1]
	}
	for _, kv := range data {
		var lower, upper int

		if key >= lower && key <= upper {
			return &kv
		}
	}
	return nil
}

func (d *Disease) GetIconFromDisease(name string) string {
	diseases := d.diseaseList
	for _, disease := range diseases {
		if disease.Name == name {
			return disease.Icon
		}
	}
	return ""
}
