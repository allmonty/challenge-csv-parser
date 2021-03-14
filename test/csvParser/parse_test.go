package test

import (
	csvparser "csvparserchallenge/src/csvParser"
	"csvparserchallenge/src/models"
	"reflect"
	"testing"
)

func TestParseWithAllNecessaryColumnsAndFilled(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"name", "email", "wage", "number"}
	inputCSV.Content = [][]string{
		{"John Doe", "doe@test.com", "$10.00", "1"},
		{"Mary Jane", "Mary@tes.com", "$15", "2"},
		{"Max Topperson", "max@test.com", "$11", "3"},
		{"Alfred Donald", "alf@test.com", "$11.5", "4"},
		{"Jane Doe", "doe@test.com", "$8.45", "5"},
	}

	result, _ := csvparser.Parse(inputCSV)

	expectedCSV := models.CSV{}
	expectedCSV.Header = []string{"email", "name", "salary"}
	expectedCSV.Content = [][]string{
		{"doe@test.com", "John Doe", "$10.00"},
		{"Mary@tes.com", "Mary Jane", "$15"},
		{"max@test.com", "Max Topperson", "$11"},
		{"alf@test.com", "Alfred Donald", "$11.5"},
		{"doe@test.com", "Jane Doe", "$8.45"},
	}

	if !reflect.DeepEqual(result, expectedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", result, expectedCSV)
	}
}

func TestParseWithAllNecessaryColumnsButOneLineWithEmailMissing(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"name", "email", "wage", "number"}
	inputCSV.Content = [][]string{
		{"John Doe", "doe@test.com", "$10.00", "1"},
		{"Alfred Donald", "", "$11.5", "4"},
	}

	parsedCSV, flaggedCSV := csvparser.Parse(inputCSV)

	expectedParsedCSV := models.CSV{}
	expectedParsedCSV.Header = []string{"email", "name", "salary"}
	expectedParsedCSV.Content = [][]string{
		{"doe@test.com", "John Doe", "$10.00"},
	}

	expectedFlaggedCSV := models.CSV{}
	expectedFlaggedCSV.Header = []string{"error", "content"}
	expectedFlaggedCSV.Content = [][]string{
		{"missing email", "Alfred Donald,,$11.5,4"},
	}

	if !reflect.DeepEqual(expectedParsedCSV, parsedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", parsedCSV, expectedParsedCSV)
	}
	if !reflect.DeepEqual(expectedFlaggedCSV, flaggedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", flaggedCSV, expectedFlaggedCSV)
	}
}

func TestParseWithoutNecessaryColumn(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"first name", "last name", "e-mail", "Rate", "Employee", "Number", "Mobile"}
	inputCSV.Content = [][]string{
		{"John", "Doe", "doe@test.com", "", "RT1", "4534151414"},
		{"Mary", "Jane", "mary@tes.com", "15", "RT2", "1448561274"},
		{"Max", "Topperson", "max@test.com", "11", "", "4534151414"},
		{"Alfred", "Donald", "alfred@test.com", "11.5", "RT4", "2145385777"},
		{"Jane", "Doe", "", "8.45", "RT5", ""},
	}

	parsedCSV, flaggedCSV := csvparser.Parse(inputCSV)

	expectedParsedCSV := models.CSV{}
	expectedParsedCSV.Header = []string{"email", "name", "salary"}

	expectedFlaggedCSV := models.CSV{}
	expectedFlaggedCSV.Header = []string{"error", "content"}
	expectedFlaggedCSV.Content = [][]string{
		{"missing salary", "John,Doe,doe@test.com,,RT1,4534151414"},
		{"missing salary", "Mary,Jane,mary@tes.com,15,RT2,1448561274"},
		{"missing salary", "Max,Topperson,max@test.com,11,,4534151414"},
		{"missing salary", "Alfred,Donald,alfred@test.com,11.5,RT4,2145385777"},
		{"missing email", "Jane,Doe,,8.45,RT5,"},
	}

	if !reflect.DeepEqual(expectedParsedCSV, parsedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", parsedCSV, expectedParsedCSV)
	}
	if !reflect.DeepEqual(expectedFlaggedCSV, flaggedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", flaggedCSV, expectedFlaggedCSV)
	}
}
