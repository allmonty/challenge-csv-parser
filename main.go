package main

import (
	csvwriter "csvparserchallenge/src/csvWriter"
	"csvparserchallenge/src/csvparser"
	"csvparserchallenge/src/csvreader"
	"csvparserchallenge/src/models"
	"fmt"
)

func main() {
	csv1 := csvreader.ReadCsvFile("examples/roster1.csv")
	csv2 := csvreader.ReadCsvFile("examples/roster2.csv")
	csv3 := csvreader.ReadCsvFile("examples/roster3.csv")
	csv4 := csvreader.ReadCsvFile("examples/roster4.csv")

	fmt.Println("CSV1=====")
	fmt.Println(csv1)
	resultCSV1, errorCSV1 := csvparser.Parse(csv1)
	fmt.Printf("\n\n%v\n\n", resultCSV1)

	fmt.Println("CSV2=====")
	fmt.Println(csv2)
	resultCSV2, errorCSV2 := csvparser.Parse(csv2)
	fmt.Printf("\n\n%v\n\n", resultCSV2)

	fmt.Println("CSV3=====")
	fmt.Println(csv3)
	resultCSV3, errorCSV3 := csvparser.Parse(csv3)
	fmt.Printf("\n\n%v\n\n", resultCSV3)

	fmt.Println("CSV4=====")
	fmt.Println(csv4)
	resultCSV4, errorCSV4 := csvparser.Parse(csv4)
	fmt.Printf("\n\n%v\n\n", resultCSV4)

	parsedCSVs := []models.CSV{
		resultCSV1,
		resultCSV2,
		resultCSV3,
		resultCSV4,
	}

	errorCSVs := []models.CSV{
		errorCSV1,
		errorCSV2,
		errorCSV3,
		errorCSV4,
	}

	csvwriter.WriteCsvFile(parsedCSVs, "result/parsed_candidates")
	csvwriter.WriteCsvFile(errorCSVs, "result/error_candidates")
}
