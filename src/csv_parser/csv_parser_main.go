package csvparser

import (
	csvreader "csvparserchallenge/src/csv_reader"
	csvwriter "csvparserchallenge/src/csv_writer"
	"csvparserchallenge/src/models"
	"fmt"
)

type processedCSV struct {
	ParsedCSV models.CSV
	ErrorCSV  models.CSV
}

func processFile(fileName string, c chan processedCSV) {
	fmt.Printf("=== Processing CSV: %v ===\n", fileName)
	csv := csvreader.ReadCsvFile(fileName)
	parsedCSV, errorCSV := Parse(csv)
	processed := processedCSV{
		ParsedCSV: parsedCSV,
		ErrorCSV:  errorCSV,
	}
	c <- processed
}

func CSVParserMain(args []string) {
	filePaths := args[1:]
	var parsedCSVs []models.CSV
	var errorCSVs []models.CSV

	channel := make(chan processedCSV)

	for _, fileName := range filePaths {
		go processFile(fileName, channel)
	}

	for range filePaths {
		processed := <-channel
		parsedCSVs = append(parsedCSVs, processed.ParsedCSV)
		errorCSVs = append(errorCSVs, processed.ErrorCSV)
	}
	close(channel)

	parsedCSV := models.CSV{Header: parsedCSVs[0].Header}
	for _, csv := range parsedCSVs {
		parsedCSV.Content = append(parsedCSV.Content, csv.Content...)
	}
	parsedCSV = RemoveDuplicatedLine(parsedCSV, "email")

	errorCSV := models.CSV{Header: errorCSVs[0].Header}
	for _, csv := range errorCSVs {
		errorCSV.Content = append(errorCSV.Content, csv.Content...)
	}

	csvwriter.WriteCsvFile(parsedCSV, "result/parsed_candidates")
	csvwriter.WriteCsvFile(errorCSV, "result/error_candidates")
}
