package test

import (
	csvparser "csvparserchallenge/src/csvParser"
	"csvparserchallenge/src/models"
	"reflect"
	"testing"
)

func TestRemoveDuplicatedLine(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"email", "name", "wage", "number"}
	inputCSV.Content = [][]string{
		{"doe@test.com", "John Doe", "$10.00", "1"},
		{"Mary@tes.com", "Mary Jane", "$15", "2"},
		{"Mary@tes.com", "Mary Jane 2", "$16", "3"},
		{"Mary@tes.com", "Mary Jane 3", "$17", "4"},
	}

	result := csvparser.RemoveDuplicatedLine(inputCSV, "email")

	expectedCSV := models.CSV{}
	expectedCSV.Header = []string{"email", "name", "wage", "number"}
	expectedCSV.Content = [][]string{
		{"doe@test.com", "John Doe", "$10.00", "1"},
		{"Mary@tes.com", "Mary Jane", "$15", "2"},
	}

	if !reflect.DeepEqual(result, expectedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", result, expectedCSV)
	}
}

func TestRemoveDuplicatedLineForEmptyContent(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"email", "name", "wage", "number"}

	result := csvparser.RemoveDuplicatedLine(inputCSV, "email")

	expectedCSV := models.CSV{}
	expectedCSV.Header = []string{"email", "name", "wage", "number"}

	if !reflect.DeepEqual(result.Header, expectedCSV.Header) {
		t.Errorf("No matching result, got: %v, want: %v.", result.Header, expectedCSV.Header)
	}

	if len(result.Content) != 0 {
		t.Errorf("Expected the result CSV with no content, got %v", expectedCSV)
	}
}
