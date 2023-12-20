package module

import (
	"gorm.io/gorm"
)

type DiseaseInfo struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Inputs      []Input `json:"inputs"`
	Output      struct {
		Description string     `json:"description"`
		Result      []KeyValue `json:"result"`
	} `json:"output"`
}

type Input struct {
	Name               string     `json:"name"`
	Description        string     `json:"description"`
	Index              int        `json:"index"`
	InputMax           float64    `json:"input_max"`
	InputMin           float64    `json:"input_min"`
	InputDecimalLength int        `json:"input_decimal_length"`
	Warn               string     `json:"warn"`
	InputMethod        string     `json:"input_method"`
	Select             []KeyValue `json:"select"`
}

type InputResult struct {
	Name               string     `json:"name"`
	Description        string     `json:"description"`
	Index              int        `json:"index"`
	InputMax           float64    `json:"input_max"`
	InputMin           float64    `json:"input_min"`
	InputDecimalLength int        `json:"input_decimal_length"`
	Warn               string     `json:"warn"`
	InputMethod        string     `json:"input_method"`
	Select             []KeyValue `json:"select"`
	Default            *KeyValue  `json:"default"`
}

type DiseaseInfoResult struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	InputResult []InputResult `json:"inputs"`
}

type KeyValue struct {
	Key   any `json:"key"`
	Value any `json:"value"`
}

type Disease struct {
	Name        string   `json:"name"`
	Index       int      `json:"index"`
	Description string   `json:"description"`
	Positions   []string `json:"positions"`
	Icon        string   `json:"icon"`
}

type DiseaseList map[string]DiseaseInfo

type Module struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserMedicalInformation struct {
	Name   string `json:"name"`
	Inputs []any  `json:"inputs"`
}

type PredictionPremise struct {
	User   string `json:"user"`
	Name   string `json:"name"`
	Module string `json:"module"`
	Inputs []struct {
		Name        string `json:"name"`
		Index       int    `json:"index"`
		SelectKey   string `json:"select-key"`
		SelectValue string `json:"select-value"`
	} `json:"inputs"`
}

type DiseasePropose struct {
	gorm.Model
	Name string `json:"name" gorm:"notnull, index:idx_name"`
	Hash string `json:"hash" gorm:"notnull"`
}

type Quantized struct {
	Data      []string
	Scale     string
	ZeroPoint string
}

type RecordPrediction struct {
	gorm.Model
	Status        Status `json:"status"`
	Message       string `json:"message"`
	User          string `json:"user"`
	Disease       string `json:"disease"`
	Icon          string `json:"icon" gorm:"type:text"`
	Module        string `json:"module"`
	Inputs        string `json:"inputs" gorm:"type:text" `
	InputsMapping string `json:"inputs-mapping" gorm:"type:text" `
	Normalized    string `json:"normalized" gorm:"type:text"`
	Quantized     string `json:"quantized" gorm:"type:text"`
	Scale         string `json:"scale"`
	ZeroPoint     string `json:"zero-point"`
	StartTime     int64  `gorm:"start_time" json:"start-time"`
	EndTime       int64  `gorm:"end_time" json:"end-time"`
	Output        int    `json:"output"`
}

type Status string

var (
	Start                Status = "start"
	NormalizationFailed  Status = "normalization failed"
	NormalizationSuccess Status = "normalization success"
	CreateSuccess        Status = "create project success"
	CreateFailed         Status = "create project failed"
	PredictionComplete   Status = "prediction complete"
	PredictionFailed     Status = "prediction failed"
	Complete             Status = "complete"
)

type GetPredictedResult struct {
	User string
	Id   uint
}

type PredictedResult struct {
	ID        uint   `json:"id"`
	Disease   string `json:"disease"`
	Icon      string `json:"icon" gorm:"type:text"`
	Module    string `json:"module"`
	StartTime int64  `gorm:"start_time" json:"start_time"`
	EndTime   int64  `gorm:"end_time" json:"end_time"`
	Status    Status `json:"status"`
	Inputs    string `json:"inputs"`
	Message   string `json:"message"`
	Output    string `json:"output"`
}

type UserPredictionValidation struct {
	gorm.Model
	User               string `gorm:"user"`
	RecordPredictionID uint   `gorm:"record_prediction_id"`
	Result             string `gorm:"result"`
	Proof              string `gorm:"type:text"`
	Disease            string `json:"disease"`
	Module             string `json:"module"`
	IsVerified         bool   `json:"is_verified"`
}

type User struct {
	gorm.Model
	Name string `gorm:"name,not null"`
}

type VerifyReq struct {
	User string `json:"user"`
	ID   uint   `json:"id"`
}

type VerifyResultNoModel struct {
	ContractAddress  string `json:"contract_address"`
	ContractFunction string `json:"contract_function"`
	Proof            string `json:"proof"`
	Result           string `json:"result"`
	Disease          string `json:"disease"`
	Module           string `json:"module"`
	ABI              string `json:"abi"`
}

type Article struct {
	gorm.Model
	Disease string `gorm:"disease"`
	URL     string `gorm:"url"`
	Time    int64  `gorm:"time"`
	Icon    string `json:"icon" gorm:"type:longtext"`
	Title   string `gorm:"title"`
}

type ArticleResult struct {
	Disease string `json:"disease" gorm:"disease"`
	URL     string `json:"url" gorm:"url"`
	Time    int64  `json:"time" gorm:"time"`
	Icon    string `json:"icon" gorm:"type:text"`
	Title   string `json:"title" gorm:"title"`
}

type ArticleCollection struct {
	gorm.Model
	User string
	Url  string
}

type InitialDiagnosisCache struct {
	gorm.Model
	Cache string `gorm:"column:cache;type:text"`
}

// InitialDiagnosisInfo DiseasesRequest Used to receive requests for disease classification
type InitialDiagnosisInfo struct {
	ChosenUser       string           `json:"chosen_user"`
	Gender           string           `json:"gender"`
	Age              int              `json:"age"`
	MedicalCondition MedicalCondition `json:"medical_condition"`
	SymptomList      []SymptomList    `json:"symptom_list"`
}

type MedicalCondition struct {
	Diabetes     string `json:"diabetes"`
	Obese        string `json:"obese"`
	Hypertension string `json:"hypertension"`
	Smoking      string `json:"smoking"`
	Anemia       string `json:"anemia"`
	Cholesterol  string `json:"cholesterol"`
}

type SymptomList struct {
	Position string   `json:"position"`
	Symptoms []string `json:"symptoms"`
}
