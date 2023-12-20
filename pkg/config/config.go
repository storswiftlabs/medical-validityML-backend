package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var (
	conf     *viper.Viper
	rootPath string
)

func init() {
	conf = viper.New()
	_, filename, _, _ := runtime.Caller(0)
	rootPath = path.Dir(path.Dir(path.Dir(filename)))
}
func NewConfig() *viper.Viper {
	var config string

	config = fmt.Sprintf("%s/config.yaml", rootPath)

	fmt.Println("load config file: ", config)

	return getConfig(config)
}

func getConfig(path string) *viper.Viper {
	conf.SetConfigFile(path)
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	return conf
}

func GetConfig() *viper.Viper {
	return conf
}
