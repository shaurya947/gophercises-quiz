package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	question, answer string
}

func main() {
	filenamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimitPtr := flag.Int("timeLimit", 30, "allowed time limit (in seconds) for completing the quiz")
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

	problems := parseCSVRecords(records)

	fmt.Printf("You will have %d seconds to finish the quiz. Press ENTER to begin.", *timeLimitPtr)
	fmt.Scanln()

	correct := 0
	go presentQuiz(problems, &correct)

	time.Sleep(time.Duration(*timeLimitPtr) * time.Second)
	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
}

func parseCSVRecords(records [][]string) []problem {
	problems := make([]problem, len(records))
	for index, record := range records {
		problems[index] = problem{question: record[0], answer: record[1]}
	}
	return problems
}

func presentQuiz(problems []problem, numCorrectPointer *int) {
	var response string
	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", index+1, problem.question)
		fmt.Scanln(&response)

		if strings.EqualFold(problem.answer, strings.TrimSpace(response)) {
			(*numCorrectPointer)++
		}
	}
}
