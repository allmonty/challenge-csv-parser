package csvwriter

import (
	"csvparserchallenge/src/models"
	"encoding/csv"
	"log"
	"os"
)

// TODO: Test this function
/*
ReadCsvFile - Given an array of CSV, write them to csv file
*/
func WriteCsvFile(csvs []models.CSV, fileName string) {
	records := [][]string{
		csvs[0].Header,
	}

	for _, csv := range csvs {
		records = append(records, csv.Content...)
	}

	file, error := os.Create(fileName + ".csv")
	defer file.Close()
	if error != nil {
		log.Fatalln("failed to open file", error)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if error := writer.Write(record); error != nil {
			log.Fatalln("error writing record to file", error)
		}
	}
}
