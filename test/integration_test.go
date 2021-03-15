package test

import (
	csvparser "csvparserchallenge/src/csv_parser"
	csvreader "csvparserchallenge/src/csv_reader"
	"csvparserchallenge/src/models"
	"log"
	"os"
	"reflect"
	"sort"
	"testing"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func sortContent(content [][]string, sortIndex int) {
	sort.SliceStable(content, func(i, j int) bool {
		return content[i][sortIndex] < content[j][sortIndex]
	})
}

func TestCSVParserMain(t *testing.T) {
	args := []string{
		"main.go",
		"./examples/roster1.csv",
		"./examples/roster2.csv",
		"./examples/roster3.csv",
		"./examples/roster4.csv",
	}

	e1 := os.Remove("./result/error_candidates.csv")
	check(e1)
	e2 := os.Remove("./result/parsed_candidates.csv")
	check(e2)

	csvparser.CSVParserMain(args)

	errorCSV := csvreader.ReadCsvFile("./result/error_candidates.csv")
	parsedCSV := csvreader.ReadCsvFile("./result/parsed_candidates.csv")

	expectedParsedCSV := models.CSV{}
	expectedParsedCSV.Header = []string{"email", "name", "salary", "id"}
	expectedParsedCSV.Content = [][]string{
		{"doe@test.com", "John Doe", "10", "RT1"},
		{"mary@tes.com ", "Mary Jane", "15", "RT2"},
		{"max@test.com", "Max Topperson", "11", "RT3"},
		{"alfred@test.com", "Alfred Donald", "11.5", "RT4"},
		{"jane.doe@test.com", "Jane Doe", "8.45", "RT5"},
		{"mary@tes.com", "Mary Jane", "15", "RT2"},
		{"maxtest.com", "Max Topperson", "11", "RT3"},
		{"matthew.doe@test.com", "Matthew Doe", "2,451.45", "RT6"},
	}

	expectedErrorCSV := models.CSV{}
	expectedErrorCSV.Header = []string{"error", "content"}
	expectedErrorCSV.Content = [][]string{
		{"missing id", "John Doe,doe@test.com,$10.00,1"},
		{"missing id", "Mary Jane,Mary@tes.com,$15,2"},
		{"missing id", "Max Topperson,max@test.com,$11,3"},
		{"missing email", "Alfred Donald,,$11.5,4"},
		{"missing id", "Jane Doe,doe@test.com,$8.45,5"},
		{"missing salary", "John,Doe,doe@test.com,,RT1,4534151414"},
		{"missing salary", "Mary,Jane,mary@tes.com,15,RT2,1448561274"},
		{"missing salary", "Max,Topperson,max@test.com,11,,4534151414"},
		{"missing salary", "Alfred,Donald,alfred@test.com,11.5,RT4,2145385777"},
		{"missing email", "Jane,Doe,,8.45,RT5,"},
		{"missing salary", "John,Doe,doe@test.com,,RT1,453 415 1414"},
	}

	sortContent(expectedParsedCSV.Content, 0)
	sortContent(parsedCSV.Content, 0)
	sortContent(expectedErrorCSV.Content, 1)
	sortContent(errorCSV.Content, 1)

	if !reflect.DeepEqual(expectedParsedCSV, parsedCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", parsedCSV, expectedParsedCSV)
	}

	if !reflect.DeepEqual(expectedErrorCSV, errorCSV) {
		t.Errorf("No matching result, got: %v, want: %v.", errorCSV, expectedErrorCSV)
	}
}
