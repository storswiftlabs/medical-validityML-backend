package python

import (
	"testing"
)

func TestQuantized(t *testing.T) {
	disease := "Acute_Inflammations"
	data := []string{"0.0", "0.0", "1.0", "0.0", "0.0", "0.0"}

	quantized, err := QuantitativeData(disease, data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Error(quantized)
}

func TestNormalizedData(t *testing.T) {
	disease := "Acute_Inflammations"
	data := []string{"35.5", "0", "1", "0", "0"}

	normailzed, err := NormalizedData(disease, data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(normailzed)
}

func TestRootPath(t *testing.T) {
	t.Error(rootPath)
}
