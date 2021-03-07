package main

import (
	"csvparserchallenge/src/csvreader"
	"csvparserchallenge/src/sum"
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(sum.Sum(1, 2))
	fmt.Println(csvreader.ReadCsvFile("examples/roster1.csv"))
}
