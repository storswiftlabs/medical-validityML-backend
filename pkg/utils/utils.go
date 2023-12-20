package utils

import (
	"encoding/json"
	"math/big"
	m "medical-zkml-backend/internal/module"
	"reflect"
	"strconv"
	"strings"
)

func Hexadecimal2Decimal(result string) int {
	bigInt := new(big.Int)
	bigInt.SetString(result, 0)

	return int(bigInt.Int64())
}

func CoverInput(premise *m.PredictionPremise, choose string) string {
	inputs := premise.Inputs

	var records []m.KeyValue
	//var coverInput []string
	if choose == "key" {
		for _, input := range inputs {
			records = append(records, m.KeyValue{
				Key:   input.Name,
				Value: input.SelectKey,
			})
		}
	} else {
		for _, input := range inputs {
			records = append(records, m.KeyValue{
				Key:   input.Name,
				Value: input.SelectValue,
			})
		}
	}
	jsonBytes, _ := json.Marshal(records)
	return string(jsonBytes)
}

func Strings2Floats(data []string) []float64 {
	floats := make([]float64, 0, len(data))

	for _, str := range data {
		num, _ := strconv.ParseFloat(str, 32)
		floats = append(floats, num)
	}

	return floats
}
func Ints2String(ints []int) string {
	strs := make([]string, 0, len(ints))

	for _, num := range ints {
		strs = append(strs, strconv.Itoa(num))
	}

	return strings.Join(strs, ",")
}

func Interface2string(v any) string {
	if v == nil {
		return ""
	} else {
		return v.(string)
	}
}

func Interface2strings(v any) []string {
	if v == nil {
		return nil
	}
	var result []string
	for _, ele := range v.([]any) {
		result = append(result, ele.(string))
	}
	return result
}

func Interface2float64(v any) float64 {
	if v == nil {
		return 0
	}

	switch reflect.TypeOf(v).String() {
	case "float64":
		return v.(float64)
	case "int":
		return float64(v.(int))
	}
	return 0.0
}

func Interface2int(v any) int {
	if v == nil {
		return 0
	}
	return v.(int)
}
