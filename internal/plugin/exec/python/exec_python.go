package python

import (
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"medical-zkml-backend/internal/module"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

var (
	python   string
	rootPath string
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	rootPath = path.Dir(filename)

}

func init() {
	if runtime.GOOS == "windows" {
		python = "py.exe"
	} else {
		python = "python3"
	}
}

// QuantitativeData Calling local Python to perform data quantization
func QuantitativeData(disease string, data []string) (*module.Quantized, error) {
	cmd := exec.Command(python, append([]string{fmt.Sprintf("%s/quantification_data.py", rootPath)}, append([]string{disease}, data...)...)...)
	// Execute commands and capture output
	output, err := cmd.Output()
	if err != nil {
		zap.L().Error("Data processing: quantification data execution failed", zap.Error(err))
		return &module.Quantized{}, err
	}

	quantized := new(module.Quantized)

	lines := strings.Split(string(output), "\n")

	quantized.Data = strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), ", ")
	quantized.Scale = strings.TrimSpace(strings.Split(lines[1], ":")[1])
	quantized.ZeroPoint = strings.TrimSpace(strings.Split(lines[2], ":")[1])
	zap.L().Info("Data processing: quantification data execution success")
	return quantized, nil
}

func NormalizedData(disease string, data []string) ([]string, error) {

	fmt.Printf("%s %v %v %v\n", python, fmt.Sprintf("%s/normalized_data.py", rootPath), disease, data)
	cmd := exec.Command(python, append([]string{fmt.Sprintf("%s/normalized_data.py", rootPath)}, append([]string{disease}, data...)...)...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		zap.L().Error("Data processing: normalization data execution failed", zap.Error(errors.New(stderr.String())))
		return nil, err
	}
	zap.L().Info("Data processing: normalization data execution success")

	return strings.Split(strings.TrimSpace(string(output[1:len(output)-2])), ", "), nil
}
