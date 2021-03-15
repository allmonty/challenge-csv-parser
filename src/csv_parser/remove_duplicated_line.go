package csvparser

import (
	"csvparserchallenge/src/models"
)

/*
RemoveDuplicatedLine - Given a CSV removes every line from content that is duplicated base on key (column)
*/
func RemoveDuplicatedLine(csv models.CSV, key string) models.CSV {
	knownLines := map[string]bool{}
	filteredList := [][]string{}
	keyIndex := FindColumn(csv, map[string]bool{key: true})

	for _, item := range csv.Content {
		indexedValue := item[keyIndex]
		if !knownLines[indexedValue] {
			knownLines[indexedValue] = true
			filteredList = append(filteredList, item)
		}
	}

	csv.Content = filteredList

	return csv
}
