package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"medical-zkml-backend/pkg/config"
	"testing"
)

var conf *viper.Viper

func init() {
	conf = config.NewConfig()
}

func TestNewLog(t *testing.T) {
	if err := InitLogger(conf); err != nil {
		t.Fatalf(err.Error())
	}
	zap.L().Info("Created successfully")
}
