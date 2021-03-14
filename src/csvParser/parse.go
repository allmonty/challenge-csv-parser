package csvparser

import (
	"csvparserchallenge/src/models"
	"strings"
)

func getColumnOrEmpty(csv models.CSV, column int) []string {
	result := GetColumn(csv, column)
	if len(result) == 0 {
		result = make([]string, len(csv.Content))
	}
	return result
}

/*
Parse -
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

	parsedCSV := models.CSV{
		Header: []string{"email", "name", "salary"},
	}
	errorCSV := models.CSV{
		Header: []string{"error", "content"},
	}

	csv = NormalizeHeader(csv)

	result := map[string][]string{}
	emailColumn := FindColumn(csv, emailPossibilities)
	nameColumn := FindColumn(csv, namePossibilities)
	firstNameColumn := FindColumn(csv, firstNamePossibilities)
	lastNameColumn := FindColumn(csv, lastNamePossibilities)
	salaryColumn := FindColumn(csv, salaryPossibilities)

	result["email"] = getColumnOrEmpty(csv, emailColumn)
	result["name"] = GetColumn(csv, nameColumn)
	result["firstname"] = GetColumn(csv, firstNameColumn)
	result["lastname"] = GetColumn(csv, lastNameColumn)
	result["salary"] = getColumnOrEmpty(csv, salaryColumn)

	if len(result["name"]) == 0 {
		for i := 0; i < len(csv.Content); i++ {
			line := result["firstname"][i] + " " + result["lastname"][i]
			result["name"] = append(result["name"], line)
		}
	}

	for i, line := range csv.Content {
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
			errorLine := []string{
				errorMessage,
				strings.Join(line[:], ","),
			}
			errorCSV.Content = append(errorCSV.Content, errorLine)
		} else {
			parsedCSV.Content = append(parsedCSV.Content, csvLine)
		}
	}

	return parsedCSV, errorCSV
}
