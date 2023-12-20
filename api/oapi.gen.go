// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// GetApiDiseasesDiseaseInfoParams defines parameters for GetApiDiseasesDiseaseInfo.
type GetApiDiseasesDiseaseInfoParams struct {
	Disease string `form:"disease" json:"disease"`
}

// PostApiDiseasesDiseaseInfoJSONBody defines parameters for PostApiDiseasesDiseaseInfo.
type PostApiDiseasesDiseaseInfoJSONBody struct {
	Disease string `json:"disease"`
	Id      int    `json:"id"`
}

// PostApiDiseasesDiseasesJSONBody defines parameters for PostApiDiseasesDiseases.
type PostApiDiseasesDiseasesJSONBody struct {
	Age              int    `json:"age"`
	ChosenUser       string `json:"chosen_user"`
	Gender           string `json:"gender"`
	MedicalCondition struct {
		Cholesterol  string `json:"cholesterol"`
		Diabetes     string `json:"diabetes"`
		Hypertensive string `json:"hypertensive"`
		Injured      string `json:"injured"`
		Obese        string `json:"obese"`
		Smoking      string `json:"smoking"`
	} `json:"medical_condition"`
	SymptomList []struct {
		Position string   `json:"position"`
		Symptoms []string `json:"symptoms"`
	} `json:"symptom_list"`
}

// PostApiOperatorDeletePredictedRecordJSONBody defines parameters for PostApiOperatorDeletePredictedRecord.
type PostApiOperatorDeletePredictedRecordJSONBody struct {
	Ids  []int  `json:"ids"`
	User string `json:"user"`
}

// PostApiOperatorArticleJSONBody defines parameters for PostApiOperatorArticle.
type PostApiOperatorArticleJSONBody struct {
	Diseases []string `json:"diseases"`
	User     string   `json:"user"`
}

// PostApiOperatorDiseasePredictionJSONBody defines parameters for PostApiOperatorDiseasePrediction.
type PostApiOperatorDiseasePredictionJSONBody struct {
	Inputs []struct {
		Index       int    `json:"index"`
		Name        string `json:"name"`
		SelectKey   string `json:"select-key"`
		SelectValue string `json:"select-value"`
	} `json:"inputs"`
	Module string `json:"module"`
	Name   string `json:"name"`
	User   string `json:"user"`
}

// PostApiOperatorDiseasePredictionReloadJSONBody defines parameters for PostApiOperatorDiseasePredictionReload.
type PostApiOperatorDiseasePredictionReloadJSONBody struct {
	Id   int    `json:"id"`
	User string `json:"user"`
}

// PostApiOperatorPredictingOutcomesJSONBody defines parameters for PostApiOperatorPredictingOutcomes.
type PostApiOperatorPredictingOutcomesJSONBody struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	User     string `json:"user"`
}

// PostApiOperatorPredictingOutcomesFuzzyQueryJSONBody defines parameters for PostApiOperatorPredictingOutcomesFuzzyQuery.
type PostApiOperatorPredictingOutcomesFuzzyQueryJSONBody struct {
	Key      string `json:"key"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	User     string `json:"user"`
}

// PostApiOperatorRecommendJSONBody defines parameters for PostApiOperatorRecommend.
type PostApiOperatorRecommendJSONBody struct {
	Disease string `json:"disease"`
}

// PostApiOperatorVerifyPredictionResultsJSONBody defines parameters for PostApiOperatorVerifyPredictionResults.
type PostApiOperatorVerifyPredictionResultsJSONBody struct {
	Id   int    `json:"id"`
	User string `json:"user"`
}

// PostApiUserArticleCollectionJSONBody defines parameters for PostApiUserArticleCollection.
type PostApiUserArticleCollectionJSONBody struct {
	Url  string `json:"url"`
	User string `json:"user"`
}

// PostApiUserArticleCollectionCheckJSONBody defines parameters for PostApiUserArticleCollectionCheck.
type PostApiUserArticleCollectionCheckJSONBody struct {
	User string `json:"user"`
}

// PostApiUserGetUserJSONBody defines parameters for PostApiUserGetUser.
type PostApiUserGetUserJSONBody struct {
	User string `json:"user"`
}

// PostApiUserRegisterJSONBody defines parameters for PostApiUserRegister.
type PostApiUserRegisterJSONBody struct {
	Address string `json:"address"`
}

// PostApiDiseasesDiseaseInfoJSONRequestBody defines body for PostApiDiseasesDiseaseInfo for application/json ContentType.
type PostApiDiseasesDiseaseInfoJSONRequestBody PostApiDiseasesDiseaseInfoJSONBody

// PostApiDiseasesDiseasesJSONRequestBody defines body for PostApiDiseasesDiseases for application/json ContentType.
type PostApiDiseasesDiseasesJSONRequestBody PostApiDiseasesDiseasesJSONBody

// PostApiOperatorDeletePredictedRecordJSONRequestBody defines body for PostApiOperatorDeletePredictedRecord for application/json ContentType.
type PostApiOperatorDeletePredictedRecordJSONRequestBody PostApiOperatorDeletePredictedRecordJSONBody

// PostApiOperatorArticleJSONRequestBody defines body for PostApiOperatorArticle for application/json ContentType.
type PostApiOperatorArticleJSONRequestBody PostApiOperatorArticleJSONBody

// PostApiOperatorDiseasePredictionJSONRequestBody defines body for PostApiOperatorDiseasePrediction for application/json ContentType.
type PostApiOperatorDiseasePredictionJSONRequestBody PostApiOperatorDiseasePredictionJSONBody

// PostApiOperatorDiseasePredictionReloadJSONRequestBody defines body for PostApiOperatorDiseasePredictionReload for application/json ContentType.
type PostApiOperatorDiseasePredictionReloadJSONRequestBody PostApiOperatorDiseasePredictionReloadJSONBody

// PostApiOperatorPredictingOutcomesJSONRequestBody defines body for PostApiOperatorPredictingOutcomes for application/json ContentType.
type PostApiOperatorPredictingOutcomesJSONRequestBody PostApiOperatorPredictingOutcomesJSONBody

// PostApiOperatorPredictingOutcomesFuzzyQueryJSONRequestBody defines body for PostApiOperatorPredictingOutcomesFuzzyQuery for application/json ContentType.
type PostApiOperatorPredictingOutcomesFuzzyQueryJSONRequestBody PostApiOperatorPredictingOutcomesFuzzyQueryJSONBody

// PostApiOperatorRecommendJSONRequestBody defines body for PostApiOperatorRecommend for application/json ContentType.
type PostApiOperatorRecommendJSONRequestBody PostApiOperatorRecommendJSONBody

// PostApiOperatorVerifyPredictionResultsJSONRequestBody defines body for PostApiOperatorVerifyPredictionResults for application/json ContentType.
type PostApiOperatorVerifyPredictionResultsJSONRequestBody PostApiOperatorVerifyPredictionResultsJSONBody

// PostApiUserArticleCollectionJSONRequestBody defines body for PostApiUserArticleCollection for application/json ContentType.
type PostApiUserArticleCollectionJSONRequestBody PostApiUserArticleCollectionJSONBody

// PostApiUserArticleCollectionCheckJSONRequestBody defines body for PostApiUserArticleCollectionCheck for application/json ContentType.
type PostApiUserArticleCollectionCheckJSONRequestBody PostApiUserArticleCollectionCheckJSONBody

// PostApiUserGetUserJSONRequestBody defines body for PostApiUserGetUser for application/json ContentType.
type PostApiUserGetUserJSONRequestBody PostApiUserGetUserJSONBody

// PostApiUserRegisterJSONRequestBody defines body for PostApiUserRegister for application/json ContentType.
type PostApiUserRegisterJSONRequestBody PostApiUserRegisterJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// disease info
	// (GET /api/diseases/disease_info)
	GetApiDiseasesDiseaseInfo(c *gin.Context, params GetApiDiseasesDiseaseInfoParams)
	// disease info
	// (POST /api/diseases/disease_info)
	PostApiDiseasesDiseaseInfo(c *gin.Context)
	// operator list
	// (GET /api/diseases/disease_list)
	GetApiDiseasesDiseaseList(c *gin.Context)
	// classification of disease
	// (POST /api/diseases/diseases)
	PostApiDiseasesDiseases(c *gin.Context)
	// Delete prediction records
	// (POST /api/operator/DeletePredictedRecord)
	PostApiOperatorDeletePredictedRecord(c *gin.Context)
	// article
	// (POST /api/operator/article)
	PostApiOperatorArticle(c *gin.Context)
	// disease prediction
	// (POST /api/operator/disease_prediction)
	PostApiOperatorDiseasePrediction(c *gin.Context)
	// Re predict
	// (POST /api/operator/disease_prediction_reload)
	PostApiOperatorDiseasePredictionReload(c *gin.Context)
	// Model List
	// (GET /api/operator/module_list)
	GetApiOperatorModuleList(c *gin.Context)
	// Forecast results
	// (POST /api/operator/predicting_outcomes)
	PostApiOperatorPredictingOutcomes(c *gin.Context)
	// Fuzzy Query of Prediction Results
	// (POST /api/operator/predicting_outcomes_fuzzy_query)
	PostApiOperatorPredictingOutcomesFuzzyQuery(c *gin.Context)
	// Disease recommendations
	// (POST /api/operator/recommend)
	PostApiOperatorRecommend(c *gin.Context)
	// Verify prediction results
	// (POST /api/operator/verify_prediction_results)
	PostApiOperatorVerifyPredictionResults(c *gin.Context)
	// Article Collection
	// (POST /api/user/article_collection)
	PostApiUserArticleCollection(c *gin.Context)
	// Article Collection Check
	// (POST /api/user/article_collection_check)
	PostApiUserArticleCollectionCheck(c *gin.Context)
	// Check if the user has been registered
	// (POST /api/user/get_user)
	PostApiUserGetUser(c *gin.Context)
	// User Registration
	// (POST /api/user/register)
	PostApiUserRegister(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// GetApiDiseasesDiseaseInfo operation middleware
func (siw *ServerInterfaceWrapper) GetApiDiseasesDiseaseInfo(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiDiseasesDiseaseInfoParams

	// ------------- Required query parameter "disease" -------------
	if paramValue := c.Query("disease"); paramValue != "" {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query argument disease is required, but not found"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "disease", c.Request.URL.Query(), &params.Disease)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter disease: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiDiseasesDiseaseInfo(c, params)
}

// PostApiDiseasesDiseaseInfo operation middleware
func (siw *ServerInterfaceWrapper) PostApiDiseasesDiseaseInfo(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiDiseasesDiseaseInfo(c)
}

// GetApiDiseasesDiseaseList operation middleware
func (siw *ServerInterfaceWrapper) GetApiDiseasesDiseaseList(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiDiseasesDiseaseList(c)
}

// PostApiDiseasesDiseases operation middleware
func (siw *ServerInterfaceWrapper) PostApiDiseasesDiseases(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiDiseasesDiseases(c)
}

// PostApiOperatorDeletePredictedRecord operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorDeletePredictedRecord(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorDeletePredictedRecord(c)
}

// PostApiOperatorArticle operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorArticle(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorArticle(c)
}

// PostApiOperatorDiseasePrediction operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorDiseasePrediction(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorDiseasePrediction(c)
}

// PostApiOperatorDiseasePredictionReload operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorDiseasePredictionReload(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorDiseasePredictionReload(c)
}

// GetApiOperatorModuleList operation middleware
func (siw *ServerInterfaceWrapper) GetApiOperatorModuleList(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetApiOperatorModuleList(c)
}

// PostApiOperatorPredictingOutcomes operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorPredictingOutcomes(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorPredictingOutcomes(c)
}

// PostApiOperatorPredictingOutcomesFuzzyQuery operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorPredictingOutcomesFuzzyQuery(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorPredictingOutcomesFuzzyQuery(c)
}

// PostApiOperatorRecommend operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorRecommend(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorRecommend(c)
}

// PostApiOperatorVerifyPredictionResults operation middleware
func (siw *ServerInterfaceWrapper) PostApiOperatorVerifyPredictionResults(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiOperatorVerifyPredictionResults(c)
}

// PostApiUserArticleCollection operation middleware
func (siw *ServerInterfaceWrapper) PostApiUserArticleCollection(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiUserArticleCollection(c)
}

// PostApiUserArticleCollectionCheck operation middleware
func (siw *ServerInterfaceWrapper) PostApiUserArticleCollectionCheck(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiUserArticleCollectionCheck(c)
}

// PostApiUserGetUser operation middleware
func (siw *ServerInterfaceWrapper) PostApiUserGetUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiUserGetUser(c)
}

// PostApiUserRegister operation middleware
func (siw *ServerInterfaceWrapper) PostApiUserRegister(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostApiUserRegister(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.GET(options.BaseURL+"/api/diseases/disease_info", wrapper.GetApiDiseasesDiseaseInfo)

	router.POST(options.BaseURL+"/api/diseases/disease_info", wrapper.PostApiDiseasesDiseaseInfo)

	router.GET(options.BaseURL+"/api/diseases/disease_list", wrapper.GetApiDiseasesDiseaseList)

	router.POST(options.BaseURL+"/api/diseases/diseases", wrapper.PostApiDiseasesDiseases)

	router.POST(options.BaseURL+"/api/operator/DeletePredictedRecord", wrapper.PostApiOperatorDeletePredictedRecord)

	router.POST(options.BaseURL+"/api/operator/article", wrapper.PostApiOperatorArticle)

	router.POST(options.BaseURL+"/api/operator/disease_prediction", wrapper.PostApiOperatorDiseasePrediction)

	router.POST(options.BaseURL+"/api/operator/disease_prediction_reload", wrapper.PostApiOperatorDiseasePredictionReload)

	router.GET(options.BaseURL+"/api/operator/module_list", wrapper.GetApiOperatorModuleList)

	router.POST(options.BaseURL+"/api/operator/predicting_outcomes", wrapper.PostApiOperatorPredictingOutcomes)

	router.POST(options.BaseURL+"/api/operator/predicting_outcomes_fuzzy_query", wrapper.PostApiOperatorPredictingOutcomesFuzzyQuery)

	router.POST(options.BaseURL+"/api/operator/recommend", wrapper.PostApiOperatorRecommend)

	router.POST(options.BaseURL+"/api/operator/verify_prediction_results", wrapper.PostApiOperatorVerifyPredictionResults)

	router.POST(options.BaseURL+"/api/user/article_collection", wrapper.PostApiUserArticleCollection)

	router.POST(options.BaseURL+"/api/user/article_collection_check", wrapper.PostApiUserArticleCollectionCheck)

	router.POST(options.BaseURL+"/api/user/get_user", wrapper.PostApiUserGetUser)

	router.POST(options.BaseURL+"/api/user/register", wrapper.PostApiUserRegister)

	return router
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xd/28cx3X/VwaLFpGB42m/3hf+0lJiZLuNY1q20wIKK8ztvrsbc3dmPTNL6qQSaIA2",
	"MeIUNdDUboEijtOkMFrUThq0AWI0/Q/6V0SU/F8UM7O7t3fcI5d3S5qS9YvEu9158z7v27w38zh8ZIUs",
	"SRkFKoW1/ei4YxE6Ztb2I0sSGYO1bSUQkRDHWw8PkvjGhDzEL1kd6xC4IIxa25bTtbu2ddyxWAoUp8Ta",
	"tryu3XWsjpViOVVErZs4JTcjIgALEMUP94uJJiDVfywFjiVh9NXI2rZeBrmTkt18TP7/q2qEostxAhK4",
	"sLbvPbLgAU5Szeoe5geECkaFpXBY29a7GfCZ1bEoTtQL+dRWx+LwbkY4RNa25Bl0LBFOIcEa9yxVrwrJ",
	"CZ1Yx8f76mWRMipAg3FtW/0XMiqBas5xmsYk1LzffEcoqZRM6RGO+kdkSYL5zNq2Tt774PEPP37y89/+",
	"/nfvK1HiOAP1RoSlZiACEXKSSiPfEtM3BMrZR0QgjChknEUwAarldgjqMeMRcITHYwgloROUsENIgMou",
	"elWqYeEUcxxK4OQhRGg0Q5JDwrjooBHH0eyAUBAEdxAnExIROesgTCOUMiEzjmNEqJB4RGIiZ1301hRQ",
	"yGhEFKeIg8hiKdCYswTJKaCYCYHYGEUsxQmhsJVyFmWh4iqEOBaIUP3eiGNCu+jNWZJKlgiEY8EQoWGc",
	"RaD4pRPQ74oUIJxqfo44Ueg6CAt0BHGs/qeMbiVMMo6IEBkIFJMDxd+EEi0ckqSYcCULTSKClINQNtxF",
	"fzYlMWhe4AEOJQpxZoSc0QPKjmgHjZicIiVpSUI9HOgh4YwqcjhGYxxKxgXCHNAIYgKHECHJUBrjGcKI",
	"sxi66C0OWOr5CT1k8SEIZFxLia+D0ulMKD9TfHCc5pJXuFkCKFRe0EERQGoEhoQkSRbrwV20E8spyyZT",
	"PVizHmacA5XxDFGmPkAHAebxDEUglW0wquknmOKJthAUYooEmVAyJiHWI0mScnYI6N0MK40rZcZkDGis",
	"hEwjckiiDMcCHRE5RRU77WrvSzNp/HPRnncOgeMJoEOmwI4zGuFCisongYYzPTyCB9a2nRO6n+AH1rbb",
	"s7uOHZTfgZwyFSv0R6v8Wjn+YND1PK/0+9d2v7O3fYfdeOWhCl4CYgil5u0AlEdOH8790FYOf4S59r1Y",
	"uxtQqZwK0SwZAUcjkEcAFJk5tBhzzqzjzjLc1/ADkmRJM7jOAtxg6HZtrwlax3a7jh8swZ2SdvHms2jA",
	"hrc6vIQ2x+suqtcbdp1+E7y9oOv3e0tw45bVayYx6tWc1aCdQ/EWoNhd2/OcXhMsdte2nd5gAczWnxAp",
	"gddh+cN1oORzaCw5azVg3gTlmjFKAIuMgw7eh5gTHWNUJDpHm/6yCGzbbSwC27b7iwo1MrixMxIvtSoI",
	"NVMhCcXhJUgiWJaE6/h+Y0ksGcP23Z29dgUwtwTF1yXg7y3jd4bBoDH+obuIf2/vjXbxD90Cv+LrEvD3",
	"l/H3fM9rit+1/RK/8YHt3d12LcC1/UICirN1JKCyWyKzCOaoB0uoHWdoN9b6MPAXtf7mlCRJyzFwGBS4",
	"NW8t4R4u4Ha6nu02RD0IajHfiG61GvIGZs3WnLWE2bGXTTzo+f2myvaDOfAc8/bO3hteq8r2g6AwcsVb",
	"W8CdZeD9YePQHvTrYAetwg7Kxa0/9NsC7S67ttfvN3btvjNcNPOdlgN63xkWbq34agv0ckrn9IZ+U892",
	"vJ5/Ste7uztt4lZzFLgVazW43zpiC5j1Bov6gTIiQJWoklEco/keUFGVHzISVoWxnNx5jj9wm6c0c3//",
	"9it3W85mCjfXLF2qEBbzOs/r2s1i3qDr+04pgVe+3Z4ENGWN33CzAj5lNCYUMEfRjOJE7zMouDE8UIV9",
	"IZ0K1OUUrjcInMBppnA36FUi3d293W/W4R0RKdZTuqafK93wVefyZEJ1doZDlaWJEMeEThA8MDquQF3O",
	"1gZu4A4aRreg77uDeb66e6fWw//vV+sBNdRzpIatOgVPOUBFxVVbr01U56GvIoXF7G3L7fqeb3uNNL7V",
	"7w57/nBQiXgpBxw57RXj5RRaGCV3lyaN4ZJN+IHtDxun8D23vywLtz1ZFDPkdmFYa8vx3cXEzuv2+o4T",
	"NNpzc7q+67mDuefvum36fUHeRLucr8syAHc5zQvcvtdrmuD6fuDNc569vW+2qnxNPVe+YctSw4vZqmcg",
	"LJOKv1OHCm9KLDOBXgEcyykS5hMb60VPZKN3FKf62COLqyzbc44tM3amFWCeOpWnld1g63j/+LhjJWJi",
	"NszYgTlzOVbfzo9dUs5S4JKYI5PiLGTp2yqIUyc1803nRxaRkIh1CGjtl08IlTABXtI21lA+NrqpPM3N",
	"YhVrxkBqhhvd1QwrrGYlIi36moHlqdIykuPq2dc9Pb54e79TvM0KE3iwhVMyZg+2yIQyrs9wysnv7Vee",
	"64MncYricUkSc46VteRmfvqcbZEvLZHOgsIK9VR1URVsTnpJGaUM20F3CXzVyGilPcwd+mKGXbjyxna0",
	"Ql1XbUZL0y+qIwfbDh/1pBuaqw5HpdYu1QKXpioibp0NHVS+HjEWA6anJMoOLEOhY2JxO7yfoqoXgcW1",
	"yRyPa4YEhBkncmbozU/Qn3z4uycf/c3v//eTk+99bnUsiSdGVXm3gFWdfMziCPj87F9UOTcrn3oIhxCz",
	"VElIPeYZ3SI0f0ut1VKmYvvmTfNFN2TJzSMY3Uw5U7K46fX6nu176rFQ/2w5nuP69tDxFCUl85SJmiaH",
	"PSZWdzno9ETIWyyaNe830LEgb3LQazSX6A4mccYB3Y4J1QngXQgZj3TSF1nbjuOctRAXxOrWtKjBAjNv",
	"uSBRWy65QPL42Ez51bVoLB93j3Eea3VQpVkclyTUh1MW/xYk2iyUktgYpVgSoBI9Ql5we8t3b6PjuhPx",
	"5RQ0z/8qxLQRTbIYK/P/ixvePX9ruP+X+t97tvrR79347ne79+ytYP+lP3rpD+qy1BBiQbK1EnbPVCh+",
	"r1uk6RcVzOuhaWMItVwozgRg9AjNQHQQZVWxnK5TcySlXOpoLQioDv0MKsideapLWb1A1sT5rSwZYY5S",
	"TGg9PPdceBUS1wXV25xQQGkmpoRO0A3ljYRmLBOIAkS6iyTjhGqnfKket3cu7oVJrgvy10go1cJFGNUa",
	"EfXo/HPRnSJ0XRDeyjhVWmVjlHGQU447iMhw2kHiCOJ48RFimYxB1gshOFcI68x1CXKal9g7YSYBvUrH",
	"MU4Sbb7Cuj4VbqnKsxJ8pdS69N58f+XJ/eZ1+erKe2UtVRrIi7q7Yd1dSOycirZTWuEV1NwX56l5vX2B",
	"Iu/ya7sXJd1XW9K5Qzvo9d2BKemOOyu62mNiKr3mXe3fUiOuqoIJWabo9XNtne7Pna+7pnrcLYutlAmS",
	"vxROQcj6fsgiH50l6ZRNOE6ns8WxFMKDM4cu7GBXBk4BR2cOvMUBC4luYxrqxqELMVy3qNdPnw+46AS3",
	"p5xREqI/JRGFWb1c8ShiCVCdbJh04o0M+Cw/spEQNc0ucjXXLZhF4rHmFvnK5XQOY604Wg6/xEg6n6Nm",
	"Ddg8thqhX1aMraO+Zqz97KPH//HB4/c+evrJp9d0+8weDocDLzg71hrrvcDOmthkWw1PwNp2g44VTpkA",
	"ej8TWi6JCm8ToEZKY0hwrAwv/62p++UvyOjJpiwGIYGzuCg1IoJHIBUU83k6U8YAVJBDKL4j9J1M25n5",
	"yEYgymciYQdK7Pqj0rn5NZp8Gbr36FTkLV/ROt8l4zEJs1jOkDjCccyOFLF9FcdOBbfFkW9O9etoxAHL",
	"qRm1f0ZE0uKri0cL4qzxv0K2NY9qhbwUCKsiryExl3/Nw0Vl1NYluWbq4oZRU11xUejsvEBZMldQW2Jp",
	"TmrOSmcBclsbrhvzcco0Vy5AZ6wjVQusjD/11plnRSX9CrV2xFRL+Tx2quZfGntHu0udgS+JsR3GN+Rh",
	"7R34DbKX2lOKjY7+w3UTnk2M8cw0RZ/QhFeTEs3nqkmNGh30kEgvZ/McogWWl2h+vbI0XRE/fu/7T371",
	"xbWuiIeLWZpJxBi/uQuqZNnjKn5IiMyp47kp2+v58PrRG+RvJFLy8zpur+P29ztWnr/ZD2xnHNy6Zfc9",
	"GAy9yPP8O3eGIxtcbzj0hsOBGzquv3Nn2LfOyG408dNRoBK4ln1qRb6zZLF5SFbk2zHSKsEW4vbmHK3p",
	"HY/f++mX//TzL3/21yf/9f7Tz375+H/+oeIjhQnW+0j59Op8xOut8BHMJQmNfTbyip38/c3bA7Sg6jY9",
	"9s/vBrjQindBU295BTlF9Vk2+pMPf/Dk3396rQ09WGHoxfZoagJ6Uas1WgnM0L35yE1WgXmjxryjorZ1",
	"whwcbJljQa/fGQ7nXxZtruZrVS3P+xDO7jdYoGoOIZdoOlV67ooD/guS8Vael1+QkH/m0fQFiQWbHvE2",
	"mm+/YyUsyvQFMbsQEqE4lhzgvL3fdhOE8w5uzzjiPKc7eGvVWeSiLM6LvwtvL9CeH7ZpTlqtRgrKC9Mt",
	"sFK7V5sr9CLl2oVWoZy7fJ6WT/nOmePqG9pMv4c5cNBVzIU6GJZaCp73M0qTc17rRdhpvAjf5xAzHK2/",
	"Ft814zeqy6ztwKmEW38cYd/p+7vBrt2/PbR73jAYjW/1+wO75w76YdgL7kSOOzy7HqsPpResutouuq5/",
	"t+pSElQmOoFu4nzRuKppr8jtCmHZ1nPaybqUhFaTuuego/VUbnyGPp/X3tbatP4MOTzrna7rVh5niORF",
	"a+yq1tiLp5ZfTXPsNfzttxcNsi8aZF80yH6dG2S//MHfnnz4y+tdfAbD3nBl8Wn2OZr0xxbF5mt6xFfR",
	"HuutbI/985dvMSakvvyYRIDV+hqyZET0OpxksSRpDOgI8AGKAXMKXCBCJUMYCckZnaAwxkKQMQFe3LAr",
	"9MULCYsg/oZARXV+CCi/a3m+4OfT1157AShR6TaODjGVeKKrg6i64SkQUTNhqaab6duKAYsZkgxlVJmr",
	"VGWPuXZYAk85SHMLcYgpmmIaxfrC5yLT1dcoExFykICUuDooIUIoOWh5CjOYySkUlzJ30SvsCA6Bd1CU",
	"mauX9I0TIoRUEgO2hm3JEDsEPiZSX2qtqGqGiy8M8Q7CUpmDGiemLIsjNAKVTOq7mBXjnJmcSOqbq8tb",
	"UPQV1zzTGpQQTil5N4PyPgzFAjqaAkWZxkZkd66PxQ3ledtu81Tnyjt1z4/Zlxaqv95ttyeffvL4J++f",
	"aru9ftuH9ooIXgQmOrnPMhmypEEHbhHL98qxrxdDN9gzTHU3qdfRP9wX5CHoHYVWT2zSlR2rlUk33mNM",
	"TZPdnGSrO46nqT/Lx98m+Xnyxd+f/OSfr7UD+c0d6P44e/hwdt/8nYr1nemOovJG/scu1nYrU+fvFIaj",
	"XEr9sHXawdrYo19VIV6N45lq8HLdb9Ucz4sTnnz6yZNf//Dk4188/fxn19MhPS/ouXawakXjELIkAdr8",
	"AOxuOeIqbmhZ62KWFRewtHv3yjNtw3lr7Re/ffrZZ9d6IXFX2O0hcDKeLZ7j6j+/09iOv6MpVM9xzfgN",
	"D3LdoPXG2uftIPc6xPB/+9HTz7/3zKRT3qIXKCUUjbP3QxbH0Kyf8G0BRf/s7fmoDQw+43EFVQIRyRKN",
	"6o8jTjgRI4jjm0dTLLeI2CKVM5GtceAMbTuKAI+ctjvNNFeb9j8pIq26jCb47Pfenvz4v59+9HcVb9Hw",
	"aj0lB35lXjJo5CX3wymEB+v5yu2p+QXS9R2mVTtvbs8tWvJzY8Mn//JXJx//4jpasm/XWPIEZPkrwuda",
	"7ssg3zY8X8xUX5hae6ami8InP/705L3fnPzj548/+NfHv/nPk19/+vj7P7qW0XNYY3McJkTIhjZ3t3h5",
	"kysGooiDEK2FyJLeeaZbvNiO9ZbUnuka0Zju9bXYfm6xxyVzj5b+zK2oXldTJu6V7zTLx/vH/x8AAP//",
	"KJYnfAV4AAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
