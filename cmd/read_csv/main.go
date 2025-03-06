package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"log"
)

func main() {
	file, err := os.Open("fil.csv")

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}