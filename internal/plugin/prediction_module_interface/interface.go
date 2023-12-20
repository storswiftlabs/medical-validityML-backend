package prediction_module_interface

import (
	m "medical-zkml-backend/internal/module"
	"medical-zkml-backend/internal/plugin/exec/cairo"
)

type DiseasePrediction interface {
	DiseasePrediction(recordPrediction *m.RecordPrediction, disease, module string, inputs []int)
}

var dp DiseasePrediction

func Register(name string) {
	switch name {
	case "Local Cairo":
		local := cairo.InitCairoRunner("scarb", "internal/plugin/circuit")
		dp = DiseasePrediction(local)
	case "Docker":
		d := cairo.NewDockerClient()
		dp = DiseasePrediction(d)
	}
}

func GetPredictionModule() DiseasePrediction {
	return dp
}
