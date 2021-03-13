package csvparser

import (
	"csvparserchallenge/src/models"
)

/*
GetColumn - Get the specified column of a CSV as an array
*/
func GetColumn(csv models.CSV, column int) []string {
	if column < 0 {
		return []string{}
	}

	var columnList []string
	for _, item := range csv.Content {
		columnList = append(columnList, item[column])
	}

	return columnList
}
