package test

import (
	csvparser "csvparserchallenge/src/csv_parser"
	"csvparserchallenge/src/models"
	"reflect"
	"testing"
)

func TestParse_WithAllNecessaryColumnsAndFilled_ReturnAllRownsInParsed(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"name", "email", "wage", "number", "id"}
	inputCSV.Content = [][]string{
		{"John Doe", "doe@test.com", "$10.00", "1", "a1"},
		{"Mary Jane", "Mary@tes.com", "$15", "2", "a2"},
		{"Max Topperson", "max@test.com", "$11", "3", "a3"},
		{"Alfred Donald", "alf@test.com", "$11.5", "4", "a4"},
		{"Jane Doe", "doe@test.com", "$8.45", "5", "a5"},
	}

	result, flaggedCSV := csvparser.Parse(inputCSV)

	expectedCSV := models.CSV{}
	expectedCSV.Header = []string{"email", "name", "salary", "id"}
	expectedCSV.Content = [][]string{
		{"doe@test.com", "John Doe", "$10.00", "a1"},
		{"Mary@tes.com", "Mary Jane", "$15", "a2"},
		{"max@test.com", "Max Topperson", "$11", "a3"},
		{"alf@test.com", "Alfred Donald", "$11.5", "a4"},
		{"doe@test.com", "Jane Doe", "$8.45", "a5"},
	}
	expectedFlaggedCSV := models.CSV{}
	expectedFlaggedCSV.Header = []string{"error", "content"}

	if !reflect.DeepEqual(result, expectedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", result, expectedCSV)
	}
	if !reflect.DeepEqual(expectedFlaggedCSV, flaggedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", flaggedCSV, expectedFlaggedCSV)
	}
}

func TestParse_WithAllNecessaryColumnsButOneLineWithEmailMissing_ReturnTheOneMissingEmailInFlagged(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"name", "email", "wage", "number", "empid"}
	inputCSV.Content = [][]string{
		{"John Doe", "doe@test.com", "$10.00", "1", "a"},
		{"Alfred Donald", "", "$11.5", "4", ""},
	}

	parsedCSV, flaggedCSV := csvparser.Parse(inputCSV)

	expectedParsedCSV := models.CSV{}
	expectedParsedCSV.Header = []string{"email", "name", "salary", "id"}
	expectedParsedCSV.Content = [][]string{
		{"doe@test.com", "John Doe", "$10.00", "a"},
	}

	expectedFlaggedCSV := models.CSV{}
	expectedFlaggedCSV.Header = []string{"error", "content"}
	expectedFlaggedCSV.Content = [][]string{
		{"missing email", "Alfred Donald,,$11.5,4,"},
	}

	if !reflect.DeepEqual(expectedParsedCSV, parsedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", parsedCSV, expectedParsedCSV)
	}
	if !reflect.DeepEqual(expectedFlaggedCSV, flaggedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", flaggedCSV, expectedFlaggedCSV)
	}
}

func TestParse_WithoutNecessaryColumn_AllLinesGoesToFlagged(t *testing.T) {
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
	expectedParsedCSV.Header = []string{"email", "name", "salary", "id"}

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
