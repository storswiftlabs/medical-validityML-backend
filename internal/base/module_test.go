package base

import (
	"github.com/stretchr/testify/assert"
	"medical-zkml-backend/internal/module"
	"medical-zkml-backend/pkg/config"
	"testing"
)

func init() {
	config.NewConfig()
}
func TestGetDiseaseConfig(t *testing.T) {
	diseaseList, _ := GetDiseaseInfo(config.GetConfig())
	assert.NotEmpty(t, diseaseList)
	t.Log(diseaseList)
}

func TestInEleAnArray(t *testing.T) {
	diseases := []string{"Chronic_Kidney_Disease", "Parkinsons", "Primary_Tumor", "Liver_Disorders,", "Heart_Failure_Clinical_Records", "Heart_Disease", "Lymphography", "Acute_Inflammations", "Breast_Cancer"}
	assert.True(t, inAnArray("Chronic_Kidney_Disease", diseases))
	assert.True(t, inAnArray("1", diseases))
}

func TestInject(t *testing.T) {
	disease := "Acute_Inflammations"
	info := new(module.DiseaseInfo)

	inject(config.GetConfig().Sub("disease").Sub(disease), info)
	t.Error("Name: ", info.Name)
	t.Error("Description: ", info.Description)
	t.Error("Inputs: ", info.Inputs)
	t.Error(info.Output)

}

func TestSubList(t *testing.T) {
	disease := "Heart_Disease"
	t.Log(config.GetConfig().Sub("disease").Sub(disease).Get("description"))
	//t.Log(config.GetConfig().Sub(disease).Get("description"))
}
