package test

import (
	csvparser "csvparserchallenge/src/csvParser"
	"csvparserchallenge/src/models"
	"reflect"
	"testing"
)

func TestGetColumn(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"name", "email", "wage", "number"}
	inputCSV.Content = [][]string{
		{"John Doe", "doe@test.com", "$10.00", "1"},
		{"Mary Jane", "Mary@tes.com", "$15", "2"},
		{"Max Topperson", "max@test.com", "$11", "3"},
		{"Alfred Donald", "", "$11.5", "4"},
		{"Jane Doe", "doe@test.com", "$8.45", "5"},
	}

	column := 1

	result := csvparser.GetColumn(inputCSV, column)

	expected := []string{"doe@test.com", "Mary@tes.com", "max@test.com", "", "doe@test.com"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("No matching result, got: %v, want: %v.", result, expected)
	}
}

func TestGetColumnWhenColumnIsNegative(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"name", "email", "wage", "number"}
	inputCSV.Content = [][]string{
		{"John Doe", "doe@test.com", "$10.00", "1"},
		{"Mary Jane", "Mary@tes.com", "$15", "2"},
		{"Max Topperson", "max@test.com", "$11", "3"},
		{"Alfred Donald", "", "$11.5", "4"},
		{"Jane Doe", "doe@test.com", "$8.45", "5"},
	}

	column := -1

	result := csvparser.GetColumn(inputCSV, column)

	expected := []string{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("No matching result, got: %v, want: %v.", result, expected)
	}
}
