package test

import (
	csvparser "csvparserchallenge/src/csv_parser"
	"csvparserchallenge/src/models"
	"testing"
)

func TestFindColumnForSinglePossibility(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"name", "email", "wage", "number"}

	possibleWords := map[string]bool{
		"wage": true,
	}

	result := csvparser.FindColumn(inputCSV, possibleWords)

	expected := 2

	if result != expected {
		t.Errorf("No matching result, got: %v, want: %v.", result, expected)
	}
}

func TestFindColumnForMultiplePossibility(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"name", "email", "wage", "number"}

	possibleWords := map[string]bool{
		"email1": true,
		"email2": true,
		"email":  true,
	}

	result := csvparser.FindColumn(inputCSV, possibleWords)

	expected := 1

	if result != expected {
		t.Errorf("No matching result, got: %v, want: %v.", result, expected)
	}
}

func TestFindColumnWhenNoMatchFound(t *testing.T) {
	inputCSV := models.CSV{}
	inputCSV.Header = []string{"name", "email", "wage", "number"}

	possibleWords := map[string]bool{
		"other1": true,
		"other2": true,
		"other3": true,
	}

	result := csvparser.FindColumn(inputCSV, possibleWords)

	expected := -1

	if result != expected {
		t.Errorf("No matching result, got: %v, want: %v.", result, expected)
	}
}
