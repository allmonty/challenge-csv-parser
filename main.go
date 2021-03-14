package main

import (
	csvwriter "csvparserchallenge/src/csvWriter"
	"csvparserchallenge/src/csvparser"
	"csvparserchallenge/src/csvreader"
	"csvparserchallenge/src/models"
	"fmt"
	"os"
)

func main() {
	var parsedCSVs []models.CSV
	var errorCSVs []models.CSV

	for _, fileName := range os.Args[1:] {
		fmt.Printf("======= Processing CSV: %v =====\n\n", fileName)
		csv := csvreader.ReadCsvFile(fileName)
		parsedCSV, errorCSV := csvparser.Parse(csv)
		parsedCSVs = append(parsedCSVs, parsedCSV)
		errorCSVs = append(errorCSVs, errorCSV)
	}

	csvwriter.WriteCsvFile(parsedCSVs, "result/parsed_candidates")
	csvwriter.WriteCsvFile(errorCSVs, "result/error_candidates")
}
