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

	numCorrectChan := make(chan int)
	go presentQuiz(problems, numCorrectChan, *timeLimitPtr)

	fmt.Printf("You scored %d out of %d.\n", <-numCorrectChan, len(problems))
}

func parseCSVRecords(records [][]string) []problem {
	problems := make([]problem, len(records))
	for index, record := range records {
		problems[index] = problem{question: record[0], answer: record[1]}
	}
	return problems
}

func presentQuiz(problems []problem, numCorrectChan chan int, timeLimit int) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	numCorrect := 0
	inputChan := make(chan string)
loop:
	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", index+1, problem.question)
		go getAnswerInput(inputChan)
		select {
		case <-timer.C:
			fmt.Println()
			break loop
		case input := <-inputChan:
			if strings.EqualFold(problem.answer, strings.TrimSpace(input)) {
				numCorrect++
			}
		}
	}

	numCorrectChan <- numCorrect
}

func getAnswerInput(inputChan chan string) {
	var input string
	fmt.Scanln(&input)
	inputChan <- input
}
