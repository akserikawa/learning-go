package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Quiz - contains the quiz score
type Quiz struct {
	Score     int `json:"score"`
	Questions int `json:"questions"`
}

func (q *Quiz) getScore() int {
	return q.Score
}

func (q *Quiz) increment() {
	q.Score++
}

func (q *Quiz) totalQuestions() int {
	return q.Questions
}

func main() {

	r := readCsv("problems.csv")
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	// Quiz setup
	q := Quiz{Score: 0, Questions: len(records)}

	for _, record := range records {

		question, correctAnswer := record[0], record[1]
		input := prompt(question)

		if isCorrect(input, correctAnswer) {
			q.increment()
		}
	}

	fmt.Printf("Quiz finished!\r\nYour score: %d/%d\n", q.getScore(), q.totalQuestions())
}

func readCsv(filename string) *csv.Reader {
	quizContent, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	return csv.NewReader(strings.NewReader(string(quizContent)))
}

func prompt(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v ", question)
	input, _ := reader.ReadString('\n')

	return input
}

func isCorrect(input string, correctAnswer string) bool {
	userAnswer := strings.TrimRight(input, "\n")

	return userAnswer == correctAnswer
}
