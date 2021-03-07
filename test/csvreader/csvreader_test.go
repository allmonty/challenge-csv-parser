package sum_test

import (
	"csvparserchallenge/src/csvreader"
	"strings"
	"testing"
)

func TestReadCsvFile(t *testing.T) {
	filePath := "../examples/roster1.csv"

	result := csvreader.ReadCsvFile(filePath)

	expected := [][]string{
		{"Name", "Email", "Wage", "Number"},
		{"John Doe", "doe@test.com", "$10.00", "1"},
		{"Mary Jane", "Mary@tes.com", "$15", "2"},
		{"Max Topperson", "max@test.com", "$11", "3"},
		{"Alfred Donald", "", "$11.5", "4"},
		{"Jane Doe", "doe@test.com", "$8.45", "5"},
	}

	errorInList := false
	for i, list := range result {
		if i == 0 {
			t.Logf("Item %s", list[0])
			t.Logf("Item %v", []byte(list[0]))
			t.Logf("Expected %v", []byte(expected[0][0]))
		}

		for j, item := range list {
			if strings.Compare(strings.TrimSpace(item), expected[i][j]) != 0 {
				errorInList = true
				t.Logf("Item %d is different of %d.", len(item), len(expected[i][j]))
				t.Logf("Item %s is different of %s.", item, expected[i][j])
			}
		}
		if errorInList {
			errorInList = false
			t.Errorf("No matching result, got: %v, want: %v.", list, expected[i])
		}
	}
}

func TestReadCsvFile2(t *testing.T) {
	filePath := "../examples/roster2.csv"

	result := csvreader.ReadCsvFile(filePath)

	expected := [][]string{
		{"Last", "E-mail", "Salary", "ID", "First"},
		{"Doe", "doe@test.com", "10", "RT1", "John"},
		{"Jane", "mary@tes.com ", "15", "RT2", "Mary"},
		{"Topperson", "max@test.com", "11", "RT3", "Max"},
		{"Donald", "alfred@test.com", "11.5", "RT4", "Alfred"},
		{"Doe", "jane.doe@test.com", "8.45", "RT5", "Jane"},
	}

	errorInList := false
	for i, list := range result {
		if i == 0 {
			continue
		}
		for j, item := range list {
			if item != expected[i][j] {
				errorInList = true
				t.Logf("Item %d is different of %d.", len(item), len(expected[i][j]))
				t.Logf("Item %s is different of %s.", item, expected[i][j])
			}
		}
		if errorInList {
			errorInList = false
			t.Errorf("No matching result, got: %v, want: %v.", list, expected[i])
		}
	}
}
