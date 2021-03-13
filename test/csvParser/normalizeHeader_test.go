package test

import (
	csvparser "csvparserchallenge/src/csvParser"
	"csvparserchallenge/src/models"
	"reflect"
	"testing"
)

func TestNormalizeHeader(t *testing.T) {
	input := models.CSV{}
	input.Header = []string{"Na;3mE", "E-Ma;il", "WaGe", "Num-_Ber"}

	expected := models.CSV{}
	expected.Header = []string{"name", "email", "wage", "number"}

	result := csvparser.NormalizeHeader(input)

	if !reflect.DeepEqual(result.Header, expected.Header) {
		t.Errorf("No matching result, got: %v, want: %v.", result.Header, expected.Header)
	}
}
