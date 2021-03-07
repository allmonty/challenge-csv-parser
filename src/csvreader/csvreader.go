package csvreader

import (
	"encoding/csv"
	"log"
	"os"

	"golang.org/x/text/encoding/unicode"
)

// ReadCsvFile auto explained
func ReadCsvFile(filePath string) [][]string {
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

	return records
}
