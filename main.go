package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Reading the csv test files
func readCsv(filepath string) [][]string {
	f, error := os.Open(filepath)
	if error != nil {
		log.Fatal("Unable to read the csv file")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	recors, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse the csv file")
	}

	return recors
}
func main() {
	fmt.Println("Hello")
	csvContent := readCsv("test_data/taxi_zone_geo.csv")
	fmt.Println(csvContent)
}
