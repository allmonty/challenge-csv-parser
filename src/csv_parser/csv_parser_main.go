package csvparser

import (
	csvreader "csvparserchallenge/src/csv_reader"
	csvwriter "csvparserchallenge/src/csv_writer"
	"csvparserchallenge/src/models"
	"fmt"
)

type processedCSV struct {
	ParsedCSV  models.CSV
	FlaggedCSV models.CSV
}

func processFile(fileName string, c chan processedCSV) {
	fmt.Printf("==> Processing CSV: %v\n", fileName)
	csv := csvreader.ReadCSVFile(fileName)
	parsedCSV, flaggedCSV := Parse(csv)
	processed := processedCSV{
		ParsedCSV:  parsedCSV,
		FlaggedCSV: flaggedCSV,
	}
	c <- processed
	fmt.Printf("=====> DONE processing CSV: %v\n", fileName)
}

//CSVParserMain - Main function
//Orquestrates the read, parsing and writing of CSV files
func CSVParserMain(args []string) {
	filePaths := args[1:]

	var parsedCSV models.CSV
	var flaggedCSV models.CSV

	channel := make(chan processedCSV)
	for _, fileName := range filePaths {
		go processFile(fileName, channel)
	}

	for i := range filePaths {
		processed := <-channel

		if i == 0 {
			parsedCSV.Header = processed.ParsedCSV.Header
			flaggedCSV.Header = processed.FlaggedCSV.Header
		}

		parsedCSV.Content = append(parsedCSV.Content, processed.ParsedCSV.Content...)
		flaggedCSV.Content = append(flaggedCSV.Content, processed.FlaggedCSV.Content...)
	}
	close(channel)

	fmt.Println("==> Removing duplicates from parsed CSV")
	parsedCSV = RemoveDuplicatedLine(parsedCSV, "email")
	fmt.Println("=====> DONE Removing duplicates from parsed CSV")

	fmt.Println("==> Writing parsed CSV")
	csvwriter.WriteCSVFile(parsedCSV, "result/parsed_candidates")
	fmt.Println("=====> DONE Writing parsed CSV")
	fmt.Println("==> Writing error CSV")
	csvwriter.WriteCSVFile(flaggedCSV, "result/error_candidates")
	fmt.Println("=====> DONE Writing parsed CSV")

	fmt.Println("______________________________")

	fmt.Printf("Sucessfuly parsed %v unique rows\n", len(parsedCSV.Content))
	fmt.Printf("Found %v rows with issues\n", len(flaggedCSV.Content))
}
