package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Problem ...
type Problem struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	filename, timeLimit := parseFlagArgs()

	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", filename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	correct := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
loop:
	for _, p := range problems {
		promptQuestion(p.Question)

		answerCh := make(chan string)
		go handleAnswers(answerCh)

		select {
		case <-timer.C:
			fmt.Println()
			break loop
		case answer := <-answerCh:
			if answer == p.Answer {
				correct++
			}
		}
	}

	fmt.Printf("Quiz finished! You scored %d out of %d.\n", correct, len(problems))
}

func parseFlagArgs() (string, int) {
	filename := flag.String("filename", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	return *filename, *timeLimit
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func promptQuestion(q string) {
	fmt.Printf("%s = ", q)
}

func handleAnswers(answerCh chan string) {
	var answer string
	fmt.Scanf("%s\n", &answer)
	answerCh <- answer
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
