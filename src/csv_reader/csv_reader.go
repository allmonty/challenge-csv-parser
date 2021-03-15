package csvreader

import (
	"csvparserchallenge/src/models"
	"encoding/csv"
	"log"
	"os"

	"golang.org/x/text/encoding/unicode"
)

/*
ReadCsvFile - Reads a CSV file and outputs it as an array
input file example:
	name,email,salary
	allan,1@2.com,10

output:
	[[name, email, salary], [allan, 1@2.com, 10]]
*/
func ReadCsvFile(filePath string) models.CSV {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read file "+filePath+"\n", err)
	}
	defer file.Close()

	dec := unicode.UTF8BOM.NewDecoder() //to fix problems with encoding ï¿»
	reader := dec.Reader(file)

	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath+"\n", err)
	}

	return models.CSV{
		Header:  records[0],
		Content: records[1:],
	}
}
