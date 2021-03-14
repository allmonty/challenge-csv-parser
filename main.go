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
	namePossibilities := map[string]bool{
		"email": true,
	}
	firstNamePossibilities := map[string]bool{
		"fname":     true,
		"firstname": true,
		"first":     true,
	}
	lastNamePossibilities := map[string]bool{
		"lname":    true,
		"lastname": true,
		"last":     true,
	}
	salaryPossibilities := map[string]bool{
		"salary": true,
		"wage":   true,
	}

	fmt.Println(csv1)
	result := map[string][]string{}
	emailColumn := csvparser.FindColumn(csv1, emailPossibilities)
	nameColumn := csvparser.FindColumn(csv1, namePossibilities)
	firstNameColumn := csvparser.FindColumn(csv1, firstNamePossibilities)
	lastNameColumn := csvparser.FindColumn(csv1, lastNamePossibilities)
	salaryColumn := csvparser.FindColumn(csv1, salaryPossibilities)

	result["email"] = csvparser.GetColumn(csv1, emailColumn)
	result["name"] = csvparser.GetColumn(csv1, nameColumn)
	result["firstname"] = csvparser.GetColumn(csv1, firstNameColumn)
	result["lastname"] = csvparser.GetColumn(csv1, lastNameColumn)
	result["salary"] = csvparser.GetColumn(csv1, salaryColumn)

	fmt.Printf("Result\n\n%v\n", result)
	fmt.Println(" ")
}
