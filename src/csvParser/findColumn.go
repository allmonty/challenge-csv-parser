package csvparser

import (
	"csvparserchallenge/src/models"
)

/*
FindColumn - Given a CSV and an array of words, find the first column the the header name is on the array

possibleWords must be a map[string]bool with each key as the possibilities

*/
func FindColumn(csv models.CSV, possibleWords map[string]bool) int {
	for i, item := range csv.Header {
		if possibleWords[item] {
			return i
		}
	}
	return -1
}
