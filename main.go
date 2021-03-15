package main

import (
	csvparser "csvparserchallenge/src/csv_parser"
	"os"
)

func main() {
	csvparser.CSVParserMain(os.Args)
}
