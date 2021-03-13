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

	csv1 = csvparser.NormalizeHeader(csv1)
	csv2 = csvparser.NormalizeHeader(csv2)
	csv3 = csvparser.NormalizeHeader(csv3)
	csv4 = csvparser.NormalizeHeader(csv4)

	emailPossibilities := map[string]bool{
		"email": true,
	}
	// namePossibilities := map[string]bool{
	// 	"email": true,
	// }
	firstNamePossibilities := map[string]bool{
		"fname":     true,
		"firstname": true,
		"first":     true,
	}
	// lastNamePossibilities := map[string]bool{
	// 	"lname":    true,
	// 	"lastname": true,
	// 	"last":     true,
	// }
	// salaryPossibilities := map[string]bool{
	// 	"salary": true,
	//  "wage": true,
	// }

	fmt.Println(csv1)
	fmt.Printf("Email is on column %v\n", csvparser.FindColumn(csv2, emailPossibilities))
	fmt.Println(" ")

	fmt.Println(csv2)
	fmt.Printf("First name is on column %v\n", csvparser.FindColumn(csv2, firstNamePossibilities))
	fmt.Println(" ")

	fmt.Println(csv3)
	fmt.Printf("First name is on column %v\n", csvparser.FindColumn(csv3, firstNamePossibilities))
	fmt.Println(" ")

	fmt.Println(csv4)
	fmt.Printf("First name is on column %v\n", csvparser.FindColumn(csv4, firstNamePossibilities))
	fmt.Println(" ")
}
