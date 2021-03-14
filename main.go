package main

import (
	"csvparserchallenge/src/csvparser"
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

	fmt.Println("CSV1=====")
	fmt.Println(csv1)
	result, _ := csvparser.Parse(csv1)
	fmt.Printf("\n\n%v\n\n", result)

	fmt.Println("CSV2=====")
	fmt.Println(csv2)
	result, _ = csvparser.Parse(csv2)
	fmt.Printf("\n\n%v\n\n", result)

	fmt.Println("CSV3=====")
	fmt.Println(csv3)
	result, _ = csvparser.Parse(csv3)
	fmt.Printf("\n\n%v\n\n", result)

	fmt.Println("CSV4=====")
	fmt.Println(csv4)
	result, _ = csvparser.Parse(csv4)
	fmt.Printf("\n\n%v\n\n", result)
}
