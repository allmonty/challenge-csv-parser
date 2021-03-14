package csvparser

import (
	"csvparserchallenge/src/models"
	"strings"
)

func getColumnOrEmpty(csv models.CSV, columNamePossibilities map[string]bool) []string {
	column := FindColumn(csv, columNamePossibilities)
	result := GetColumn(csv, column)
	if len(result) == 0 {
		result = make([]string, len(csv.Content))
	}
	return result
}

func getNameColumn(csv models.CSV, namePossibilities map[string]bool, firstNamePossibilities map[string]bool, lastNamePossibilities map[string]bool) []string {
	nameColumnIndex := FindColumn(csv, namePossibilities)
	nameColumn := GetColumn(csv, nameColumnIndex)

	if len(nameColumn) == 0 {
		firstNameColumn := getColumnOrEmpty(csv, firstNamePossibilities)
		lastNameColumn := getColumnOrEmpty(csv, lastNamePossibilities)

		for i := range csv.Content {
			line := firstNameColumn[i] + " " + lastNameColumn[i]
			nameColumn = append(nameColumn, line)
		}
	}

	return nameColumn
}

/*
Parse - parse a CSV to one with defined header
returns:
- Parsed CSV with the lines that have all necessary columns
- Error CSV with error reason and content of line
*/
func Parse(csv models.CSV) (models.CSV, models.CSV) {
	emailPossibilities := map[string]bool{
		"email": true,
	}
	namePossibilities := map[string]bool{
		"name": true,
	}
	firstNamePossibilities := map[string]bool{
		"fname":     true,
		"firstname": true,
		"first":     true,
	}
	lastNamePossibilities := map[string]bool{
		"lname":    true,
		"lastname": true,
		"last":     true,
	}
	salaryPossibilities := map[string]bool{
		"salary": true,
		"wage":   true,
	}
	idPossibilities := map[string]bool{
		"id": true,
		"employeenumber": true,
		"empid": true,
	}

	parsedCSV := models.CSV{
		Header: []string{"email", "name", "salary", "id"},
	}
	errorCSV := models.CSV{
		Header: []string{"error", "content"},
	}

	csv = NormalizeHeader(csv)

	result := map[string][]string{}
	result["email"] = getColumnOrEmpty(csv, emailPossibilities)
	result["name"] = getNameColumn(csv, namePossibilities, firstNamePossibilities, lastNamePossibilities)
	result["salary"] = getColumnOrEmpty(csv, salaryPossibilities)
	result["id"] = getColumnOrEmpty(csv, idPossibilities)

	for i := range csv.Content {
		csvLine := []string{}
		error := false
		var errorMessage string
		for _, item := range parsedCSV.Header {
			if result[item][i] == "" {
				error = true
				errorMessage = "missing " + item
				break
			} else {
				csvLine = append(csvLine, result[item][i])
			}
		}

		if error {
			errorLineContent := strings.Join(csv.Content[i][:], ",")
			errorLine := []string{errorMessage, errorLineContent}
			errorCSV.Content = append(errorCSV.Content, errorLine)
		} else {
			parsedCSV.Content = append(parsedCSV.Content, csvLine)
		}
	}

	return parsedCSV, errorCSV
}
