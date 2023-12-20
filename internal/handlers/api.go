package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"medical-zkml-backend/api"
	"medical-zkml-backend/internal/module"
	"medical-zkml-backend/internal/plugin/exec/giza"
	"net/http"
	"sort"
)

type Medical struct {
	Disease
	Operator
	User
	giza.GizaRunner
}

func (m Medical) GetApiDiseasesDiseaseInfo(c *gin.Context, params api.GetApiDiseasesDiseaseInfoParams) {
	if m.Disease.diseaseList == nil {
		returnError(http.StatusInternalServerError, "", errors.New("no disease list"), c)
		return
	}

	if !m.Disease.elemInArray(params.Disease) {
		message := "there are no such diseases in the database"
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	returnDataNoCount(http.StatusOK, "", m.Disease.diseaseInfo[params.Disease], c)
	return
}

func (m Medical) PostApiDiseasesDiseaseInfo(c *gin.Context) {
	type Request struct {
		Disease string `json:"disease"`
		ID      uint   `json:"id"`
	}

	var req Request
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, err.Error(), err, c)
		return
	}

	if m.Disease.diseaseList == nil {
		returnError(http.StatusInternalServerError, "", errors.New("no disease list"), c)
		return
	}

	if !m.Disease.elemInArray(req.Disease) {
		message := "there are no such diseases in the database"
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	cache, err := m.Disease.GetInitialDiagnosisInfo(req.ID)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	returnDataNoCount(http.StatusOK, "", m.Disease.CopyDisease2Result(req.Disease, cache), c)
	return
}

func (m Medical) PostApiDiseasesDiseases(c *gin.Context) {

	var req module.InitialDiagnosisInfo
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, err.Error(), err, c)
		return
	}

	if len(req.SymptomList) == 0 {
		message := "position cannot be empty"
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	count := int64(len(m.Disease.diseaseList))
	if count == 0 {
		returnError(http.StatusInternalServerError, "", errors.New("disease list loading failed"), c)
		return
	}
	// Caching initial diagnosis information
	id, err := m.Disease.CacheInitialDiagnosisInfo(&req)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	// Search for diseases by position
	diseases := m.Disease.GetDiseaseFromPosition(&req)
	type Result struct {
		ID       uint             `json:"id"`
		Diseases []module.Disease `json:"diseases"`
	}
	returnData(http.StatusOK, int64(len(diseases)), "Query completed", Result{id, diseases}, c)
	return
}

func (m Medical) GetApiDiseasesDiseaseList(c *gin.Context) {
	count := int64(len(m.Disease.diseaseList))
	if count == 0 {
		returnError(http.StatusInternalServerError, "", errors.New("disease list loading failed"), c)
		return
	}
	list := m.Disease.diseaseList
	sort.Slice(list, func(i, j int) bool {
		return list[i].Index < list[j].Index
	})
	returnData(http.StatusOK, count, "Query completed", list, c)
	return
}

func (m Medical) GetApiOperatorModuleList(c *gin.Context) {
	count := int64(len(m.Operator.Modules))
	if count == 0 {
		returnError(http.StatusInternalServerError, "", errors.New("model list loading failed"), c)
		return
	}
	returnData(http.StatusOK, count, "ok", m.Operator.Modules, c)
	return
}

func (m Medical) PostApiOperatorDiseasePrediction(c *gin.Context) {
	var premise module.PredictionPremise

	// Obtain the input body parameter
	if err := c.BindJSON(&premise); err != nil {
		returnError(http.StatusBadRequest, "", err, c)
		return
	}

	if premise.User == "" {
		msg := "user cannot be empty"
		returnError(http.StatusBadRequest, msg, errors.New(msg), c)
		return
	}

	isExists, err := m.User.IsRegistered(premise.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}

	if !isExists {
		message := fmt.Sprintf("This user %s has not been registered yet", premise.User)
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	// Check the correctness of the data
	// Check if the disease name is in the disease list
	if !m.Disease.elemInArray(premise.Name) {
		returnError(http.StatusBadRequest, "", errors.New(fmt.Sprintf("this disease %s is not found in the lists of diseases", premise.Name)), c)
		return
	}

	// Check if the module is in the module list
	if !m.Operator.InModuleList(premise.Module) {
		returnError(http.StatusBadRequest, "", errors.New(fmt.Sprintf("this module %s is not found in the lists of modules", premise.Module)), c)
		return
	}

	// Check if the number of inputs parameters is equal
	if len(premise.Inputs) != m.Disease.CountInputsForName(premise.Name) {
		returnError(http.StatusBadRequest, fmt.Sprintf("%d%d", len(premise.Inputs), m.Disease.CountInputsForName(premise.Name)), errors.New("the number of input parameters entered does not match"), c)
		return
	}

	// Check if the index is continuous
	sort.Slice(premise.Inputs, func(i, j int) bool {
		return premise.Inputs[i].Index < premise.Inputs[j].Index
	})

	for index, input := range premise.Inputs {
		if index != input.Index {
			msg := "index error in input parameters"
			returnError(http.StatusBadRequest, msg, errors.New(msg), c)
			return
		}
	}

	if err := m.Disease.ParameterCheck(&premise); err != nil {
		returnError(http.StatusBadRequest, err.Error(), err, c)
		return
	}

	go m.Operator.DiseasePrediction(&premise)

	returnDataNoCount(http.StatusOK, "ok", nil, c)
	return
}

func (m Medical) PostApiOperatorDiseasePredictionReload(c *gin.Context) {
	// Request
	type Request struct {
		ID   uint   `json:"id"`
		User string `json:"user"`
	}
	var req Request
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, err.Error(), err, c)
		return
	}

	record, err := m.Operator.GetPredicted(req.User, req.ID)
	if err != nil {
		returnError(http.StatusBadRequest, err.Error(), err, c)
		return
	}

	disease := m.Disease.diseaseInfo[record.Disease]

	var kv []module.KeyValue
	_ = json.Unmarshal([]byte(record.InputsMapping), &kv)

	returnDataNoCount(http.StatusOK, "ok", m.Operator.Reload(&disease, kv), c)
	return
}

func (m Medical) PostApiOperatorPredictingOutcomes(c *gin.Context) {
	var req struct {
		User     string `json:"user"`
		Page     int    `json:"page"`
		PageSize int    `json:"page_size"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "", err, c)
		return
	}

	if req.User == "" {
		msg := "user cannot be empty"
		returnError(http.StatusBadRequest, msg, errors.New(msg), c)
		return
	}

	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}
	if !isExists {
		message := fmt.Sprintf("This user %s has not been registered yet", req.User)
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	// Query database to obtain results
	list, count, err := m.Operator.GetPrediction(req.User, req.Page, req.PageSize)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	for i := 0; i < len(list); i++ {
		list[i].Icon = m.Disease.GetIconFromDisease(list[i].Disease)
		if list[i].Output == "-1" {
			continue
		}
		value, err := m.Disease.GetOutputValue(list[i].Disease, list[i].Output)
		if err != nil {
			returnError(http.StatusInternalServerError, err.Error(), err, c)
			return
		}
		list[i].Output = value
	}

	returnData(http.StatusOK, count, "ok", list, c)
	return
}

func (m Medical) PostApiOperatorPredictingOutcomesFuzzyQuery(c *gin.Context) {
	type Request struct {
		User     string `json:"user"`
		Key      string `json:"key"`
		Page     int    `json:"page"`
		PageSize int    `json:"page_size"`
	}

	var req Request
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, err.Error(), err, c)
		return
	}

	if req.User == "" {
		msg := "user cannot be empty"
		returnError(http.StatusBadRequest, msg, errors.New(msg), c)
		return
	}

	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}

	if !isExists {
		message := fmt.Sprintf("This user %s has not been registered yet", req.User)
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	result, count, err := m.Operator.PredictingOutcomesFuzzyQuery(req.User, req.Key, req.Page, req.PageSize)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}

	for i := 0; i < len(result); i++ {
		result[i].Icon = m.Disease.GetIconFromDisease(result[i].Disease)
		if result[i].Output == "-1" {
			continue
		}
		value, err := m.Disease.GetOutputValue(result[i].Disease, result[i].Output)
		if err != nil {
			returnError(http.StatusInternalServerError, err.Error(), err, c)
			return
		}
		result[i].Output = value
	}

	returnData(http.StatusOK, count, "query complete", result, c)
	return
}

func (m Medical) PostApiOperatorDeletePredictedRecord(c *gin.Context) {
	var req struct {
		User string `json:"user"`
		IDs  []int  `json:"ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "", err, c)
		return
	}

	if req.User == "" {
		msg := "user cannot be empty"
		returnError(http.StatusBadRequest, msg, errors.New(msg), c)
		return
	}

	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}
	if !isExists {
		message := fmt.Sprintf("This user %s has not been registered yet", req.User)
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	if err := m.Operator.DeletePredictedResults(req.User, req.IDs); err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	returnDataNoCount(http.StatusOK, "Record deleted successfully", nil, c)
	return
}

// PostApiOperatorRecommend TODO: Accessing GPT for disease advice
func (m Medical) PostApiOperatorRecommend(c *gin.Context) {
	var req struct {
		Disease string `json:"disease"`
	}
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}

	hash, err := m.Operator.GetProposeByName(req.Disease)
	if err != nil {
		returnError(http.StatusInternalServerError, "", err, c)
		return
	}
	if hash == "" {
		returnError(http.StatusInternalServerError, "", errors.New("the suggestion list is empty"), c)
		return
	}
	returnDataNoCount(http.StatusOK, "ok", hash, c)
	return
}

func (m Medical) PostApiOperatorVerifyPredictionResults(c *gin.Context) {
	var req module.VerifyReq
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}
	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}
	if !isExists {
		message := fmt.Sprintf("This user %s has not been registered yet", req.User)
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	information, err := m.Operator.VerifyInformation(&req)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	returnDataNoCount(http.StatusOK, "ok", information, c)
	return
}

func (m Medical) PostApiOperatorArticle(c *gin.Context) {
	var req struct {
		Diseases []string `json:"diseases"`
	}
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}

	for _, disease := range req.Diseases {
		if !m.Disease.elemInArray(disease) {
			message := "there are no such diseases in the database"
			returnError(http.StatusBadRequest, message, errors.New(message), c)
			return
		}
	}

	articles, err := m.Operator.GetArticle(req.Diseases)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}
	returnData(http.StatusOK, int64(len(articles)), "ok", articles, c)
	return
}

func (m Medical) PostApiUserArticleCollection(c *gin.Context) {
	var req struct {
		User string `json:"user"`
		Url  string `json:"url"`
	}

	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, err.Error(), err, c)
		return
	}

	collect, err := m.User.IsCollected(req.User, req.Url)
	if err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}
	if collect {
		if err := m.User.CancelArticleCollection(req.User, req.Url); err != nil {
			returnError(http.StatusInternalServerError, err.Error(), err, c)
			return
		}
		returnDataNoCount(http.StatusOK, "Successfully cancelled collecting articles", nil, c)
		return
	}

	if err := m.User.CollectArticles(req.User, req.Url); err != nil {
		message := fmt.Sprintf("Article collection failed")
		returnError(http.StatusInternalServerError, message, errors.New(fmt.Sprintf("%s: %s", message, err.Error())), c)
		return
	}
	returnDataNoCount(http.StatusOK, "Successfully collected articles", nil, c)
	return
}

func (m Medical) PostApiUserRegister(c *gin.Context) {
	var req struct {
		User string `json:"address"`
	}
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}
	if req.User == "" {
		message := fmt.Sprintf("Address cannot be empty")
		returnError(http.StatusBadRequest, message, errors.New(message), c)
		return
	}

	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}

	if isExists {
		returnDataNoCount(http.StatusOK, "User already exists", nil, c)
		return
	}

	if err := m.User.UserRegistration(req.User); err != nil {
		returnError(http.StatusInternalServerError, err.Error(), err, c)
		return
	}

	returnDataNoCount(http.StatusOK, "login was successful", nil, c)
	return
}

func (m Medical) PostApiUserGetUser(c *gin.Context) {
	var req struct {
		User string `json:"user"`
	}
	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}
	isExists, err := m.User.IsRegistered(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}

	returnDataNoCount(http.StatusOK, "ok", isExists, c)
	return
}

func (m Medical) PostApiUserArticleCollectionCheck(c *gin.Context) {
	var req struct {
		User string `json:"user"`
	}

	if err := c.BindJSON(&req); err != nil {
		returnError(http.StatusBadRequest, "invalid argument", err, c)
		return
	}

	list, err := m.User.FavoriteArticleList(req.User)
	if err != nil {
		returnError(http.StatusInternalServerError, "Database query failed", err, c)
		return
	}
	returnData(http.StatusOK, int64(len(list)), "ok", list, c)
	return
}
