package csvparser

import (
	"csvparserchallenge/src/models"
	"log"
	"regexp"
	"strings"
)

/*
NormalizeHeader - Given a CSV, normalizes it's header turning words to lowercase and removing special characters
*/
func NormalizeHeader(csv models.CSV) models.CSV {
	regex, err := regexp.Compile(`[^a-zA-Z]`)
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range csv.Header {
		var word = csv.Header[i]

		word = strings.ToLower(word)
		word = regex.ReplaceAllString(word, "")

		csv.Header[i] = word
	}

	return csv
}
