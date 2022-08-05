package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Question struct {
	problem string
	answer  string
}

var points int
var totQuestions int

func main() {

	csvFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		question := Question{
			problem: line[0],
			answer:  line[1],
		}
		askQuestion(question)
		totQuestions += 1
	}
	fmt.Println("You got", points, "/", totQuestions, "questions correct!")
}

func askQuestion(question Question) {
	fmt.Println(question.problem)
	var a1 string
	fmt.Scanln(&a1)

	if a1 == question.answer {
		points += 1
	}
}
