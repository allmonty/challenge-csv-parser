package main

import (
	csvparser "csvparserchallenge/src/csvParser"
	"csvparserchallenge/src/csvreader"
	"fmt"
)

func main() {

	// parsedCSV := models.CSV{
	// 	Header: []string{"email", "data"},
	// }

	csv1 := csvreader.ReadCsvFile("examples/roster1.csv")
	csv2 := csvreader.ReadCsvFile("examples/roster2.csv")
	csv3 := csvreader.ReadCsvFile("examples/roster3.csv")
	csv4 := csvreader.ReadCsvFile("examples/roster4.csv")

	fmt.Println(csvparser.NormalizeHeader(csv1))
	fmt.Println()
	fmt.Println(csvparser.NormalizeHeader(csv2))
	fmt.Println()
	fmt.Println(csvparser.NormalizeHeader(csv3))
	fmt.Println()
	fmt.Println(csvparser.NormalizeHeader(csv4))
}
