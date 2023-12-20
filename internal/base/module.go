package base

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"image"
	"image/png"
	"medical-zkml-backend/internal/module"
	"medical-zkml-backend/pkg/utils"
	"os"
	"os/exec"
	"path"
	"reflect"
	"runtime"
	"strings"
)

var rootPath string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	rootPath = path.Dir(path.Dir(path.Dir(filename)))

}

// GetDiseaseInfo Basic data loading
func GetDiseaseInfo(conf *viper.Viper) ([]module.Disease, module.DiseaseList) {
	diseaseList := make(module.DiseaseList)
	var diseases []module.Disease
	diseaseConf := conf.Sub("disease")
	if diseaseConf == nil {
		zap.L().Panic("Failed to obtain disease configuration")
	}

	list := sublist(diseaseConf)
	if list == nil {
		zap.L().Panic("Failed to read sublist: ", zap.Error(errors.New("sublist == nil")))
	}

	for _, disease := range list {
		diseases = append(diseases, module.Disease{
			Name:        disease,
			Description: utils.Interface2string(diseaseConf.Sub(disease).Get("description")),
			Positions:   utils.Interface2strings(diseaseConf.Sub(disease).Get("positions")),
			Index:       diseaseConf.Sub(disease).Get("index").(int),
			Icon:        GetIconFromDisease(strings.ReplaceAll(disease, " ", "_")),
		})
	}

	for _, key := range list {
		info := module.DiseaseInfo{
			Name: key,
		}
		disease := diseaseConf.Sub(key)
		inject(disease, &info)
		diseaseList[key] = info
	}
	zap.L().Info("Disease list, disease information loading completed")
	return diseases, diseaseList
}

func inAnArray(ele string, arr []string) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}

func inject(disease *viper.Viper, info *module.DiseaseInfo) {

	info.Description = utils.Interface2string(disease.Get("description"))
	inputs := disease.Get("inputs").([]any)
	for _, elem := range inputs {
		param := elem.(map[string]any)
		info.Inputs = append(info.Inputs, module.Input{
			Name:               utils.Interface2string(param["name"]),
			Description:        utils.Interface2string(param["description"]),
			Index:              param["index"].(int),
			InputMax:           utils.Interface2float64(param["input_max"]),
			InputMin:           utils.Interface2float64(param["input_min"]),
			InputDecimalLength: utils.Interface2int(param["input_decimal_length"]),
			Warn:               utils.Interface2string(param["warn"]),
			InputMethod:        utils.Interface2string(param["input method"]),
			Select:             interface2keyvalue(param["select"]),
		})
	}

	output := disease.Get("output").(map[string]any)
	info.Output.Description = utils.Interface2string(output["description"])
	info.Output.Result = interface2keyvalue(output["result"])
}

// sublist Obtain a secondary list
func sublist(meet *viper.Viper) (list []string) {
	allKeys := meet.AllKeys()

	if len(allKeys) == 0 {
		return nil
	}
	for _, key := range allKeys {
		disease := cases.Title(language.Dutch).String(strings.Split(key, ".")[0])
		if len(list) == 0 {
			list = append(list, disease)
			continue
		}
		if inAnArray(disease, list) {
			continue
		}
		list = append(list, disease)
	}

	zap.L().Info("list := sublist(diseaseConf): ", zap.Any("list", list))
	return
}

func interface2keyvalue(v any) []module.KeyValue {
	if v == nil {
		panic("The selection option cannot be empty")
	}

	var keyvalues []module.KeyValue
	for _, kv := range v.([]any) {
		mapType := reflect.TypeOf(kv)
		if mapType.Kind() != reflect.Map {
			panic("The selection option must be a map")
		}
		keyType := mapType.Key()
		valueType := mapType.Elem()

		mapValue := reflect.ValueOf(kv)
		keys := mapValue.MapKeys()
		for _, key := range keys {
			keyvalues = append(keyvalues, module.KeyValue{
				Key:   key.Convert(keyType).Interface(),
				Value: mapValue.MapIndex(key).Convert(valueType).Interface(),
			})
		}

	}
	return keyvalues
}

func GetModuleList(conf *viper.Viper) []module.Module {
	var modules []module.Module
	moduleConf := conf.Get("module").([]any)

	if len(moduleConf) == 0 {
		panic("Module list cannot be empty")
	}

	for _, value := range moduleConf {
		temp := value.(map[string]any)
		modules = append(modules, module.Module{
			Name:        utils.Interface2string(temp["name"]),
			Description: utils.Interface2string(temp["description"]),
		})
	}
	return modules
}

func ExpectCheck() {
	expectCheck := exec.Command("which", "expect")
	exist, _ := expectCheck.Output()
	if string(exist) == "" {
		fmt.Println("Please use apt or yum to install the expect command")
		os.Exit(4)
	}
}

func GetIconFromDisease(name string) string {
	file, err := os.Open(fmt.Sprintf("%s/icon/%s.png", rootPath, name))
	if err != nil {
		fmt.Println("Icon loading failed: ", err)
		os.Exit(6)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Icon loading failed: ", err)
		os.Exit(6)
	}

	buf := new(bytes.Buffer)

	if err = png.Encode(buf, img); err != nil {
		fmt.Println("Icon loading failed: ", err)
		os.Exit(6)
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
