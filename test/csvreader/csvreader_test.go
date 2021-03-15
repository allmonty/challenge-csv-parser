package test

import (
	csvreader "csvparserchallenge/src/csv_reader"
	"csvparserchallenge/src/models"
	"reflect"
	"testing"
)

func TestReadCSVFile(t *testing.T) {
	filePath := "../examples/roster1.csv"

	result := csvreader.ReadCSVFile(filePath)

	expected := models.CSV{}
	expected.Header = []string{"Name", "Email", "Wage", "Number"}
	expected.Content = [][]string{
		{"John Doe", "doe@test.com", "$10.00", "1"},
		{"Mary Jane", "Mary@tes.com", "$15", "2"},
		{"Max Topperson", "max@test.com", "$11", "3"},
		{"Alfred Donald", "", "$11.5", "4"},
		{"Jane Doe", "doe@test.com", "$8.45", "5"},
	}

	if !reflect.DeepEqual(result.Header, expected.Header) {
		t.Errorf("No matching result, got: %v, want: %v.", result.Header, expected.Header)
	}

	for i, _ := range result.Content {
		if !reflect.DeepEqual(result.Content[i], expected.Content[i]) {
			t.Errorf("No matching result, got: %v, want: %v.", result.Content[i], expected.Content[i])
		}
	}
}

func TestReadCSVFile2(t *testing.T) {
	filePath := "../examples/roster2.csv"

	result := csvreader.ReadCSVFile(filePath)

	expected := models.CSV{}
	expected.Header = []string{"Last", "E-mail", "Salary", "ID", "First"}
	expected.Content = [][]string{
		{"Doe", "doe@test.com", "10", "RT1", "John"},
		{"Jane", "mary@tes.com ", "15", "RT2", "Mary"},
		{"Topperson", "max@test.com", "11", "RT3", "Max"},
		{"Donald", "alfred@test.com", "11.5", "RT4", "Alfred"},
		{"Doe", "jane.doe@test.com", "8.45", "RT5", "Jane"},
	}

	if !reflect.DeepEqual(result.Header, expected.Header) {
		t.Errorf("No matching result, got: %v, want: %v.", result.Header, expected.Header)
	}

	for i, _ := range result.Content {
		if !reflect.DeepEqual(result.Content[i], expected.Content[i]) {
			t.Errorf("No matching result, got: %v, want: %v.", result.Content[i], expected.Content[i])
		}
	}
}
