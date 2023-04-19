package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filenamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	filePtr, err := os.Open(*filenamePtr)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(filePtr)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	numRightAnswers := 0
	var response string
	for index, record := range records {
		question, answer := record[0], record[1]
		fmt.Printf("Problem #%d: %s = ", index+1, question)
		fmt.Scanln(&response)

		if strings.EqualFold(answer, strings.TrimSpace(response)) {
			numRightAnswers++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", numRightAnswers, len(records))
}
