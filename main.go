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

	parsedCSV := models.CSV{Header: parsedCSVs[0].Header}
	for _, csv := range parsedCSVs {
		parsedCSV.Content = append(parsedCSV.Content, csv.Content...)
	}
	parsedCSV = csvparser.RemoveDuplicatedLine(parsedCSV, "email")

	errorCSV := models.CSV{Header: errorCSVs[0].Header}
	for _, csv := range errorCSVs {
		errorCSV.Content = append(errorCSV.Content, csv.Content...)
	}

	csvwriter.WriteCsvFile(parsedCSV, "result/parsed_candidates")
	csvwriter.WriteCsvFile(errorCSV, "result/error_candidates")
}
