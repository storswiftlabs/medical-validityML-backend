package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	syslog "log"
	"medical-zkml-backend/api"
	"medical-zkml-backend/internal/base"
	"medical-zkml-backend/internal/db"
	"medical-zkml-backend/internal/handlers"
	"medical-zkml-backend/internal/middleware"
	"medical-zkml-backend/internal/plugin/exec/giza"
	"medical-zkml-backend/internal/plugin/prediction_module_interface"
	"medical-zkml-backend/pkg/config"
	"medical-zkml-backend/pkg/log"
)

func main() {

	// base.ExpectCheck()

	// read configuration
	conf := config.NewConfig()

	// Load Log
	if err := log.InitLogger(conf); err != nil {
		panic("Log initialization failed")
	}
	zap.L().Info("Log initialization successful")

	// Load MySQL Connection
	db.InitMysql(conf)

	// Load Disease Info
	disease := handlers.NewDisease(base.GetDiseaseInfo(conf))
	zap.L().Info("Disease module initialization successful")

	prediction_module_interface.Register("Docker")
	zap.L().Info("Cairo module initialization successful")

	// Load Module List
	modules := base.GetModuleList(conf)
	operator := handlers.Operator{
		Modules: modules,
	}

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(log.GinLogger(), log.GinRecovery(true))

	// Register disease Module
	api.RegisterHandlers(r, handlers.Medical{
		Disease:    disease,
		Operator:   operator,
		GizaRunner: giza.NewGizaRunner(conf),
	})

	//gin.DefaultWriter = io.Writer(os.Stdout)
	addr := ":3001"
	listener := "http://0.0.0.0:"

	zap.L().Info("server start", zap.String("host", listener+conf.GetString("http.port")))
	if err := r.Run(addr); err != nil {
		syslog.Panic("err:", err)
	}
}
