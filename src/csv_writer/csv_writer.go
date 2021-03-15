package csvwriter

import (
	"csvparserchallenge/src/models"
	encondingCSV "encoding/csv"
	"log"
	"os"
)

// TODO: Test this function
/*
ReadCSVFile - Given an array of CSV, write them to csv file
*/
func WriteCSVFile(csv models.CSV, fileName string) {
	var records [][]string
	records = append(records, csv.Header)
	records = append(records, csv.Content...)

	file, error := os.Create(fileName + ".csv")
	defer file.Close()
	if error != nil {
		log.Fatalln("failed to open file", error)
	}

	writer := encondingCSV.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if error := writer.Write(record); error != nil {
			log.Fatalln("error writing record to file", error)
		}
	}
}
