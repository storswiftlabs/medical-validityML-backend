package handlers

import (
	"fmt"
	"medical-zkml-backend/pkg/config"
	"os"
	"testing"
)

func TestReadJson(t *testing.T) {
	data, err := os.ReadFile("../plugin/abi/decision_tree/Lymphography.json")
	if err != nil {
		t.Errorf(err.Error())
	}
	content := string(data)
	t.Log(content)
}

func TestFunc(t *testing.T) {
	disease := "Lymphography"
	module := "decision_tree"
	config.NewConfig()
	t.Log(config.GetConfig().Get("contract"))
	t.Log(config.GetConfig().Sub("contract"))
	t.Log(config.GetConfig().Sub("contract").Sub(fmt.Sprintf("%s+%s", module, disease)).Get("address"))
}
