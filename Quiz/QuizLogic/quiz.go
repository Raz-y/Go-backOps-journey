package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Define the minimum number of columns expected in the CSV file.
const MinColumns = 2

// Represent a quiz question and its corresponding answer.
type problem struct {
	q string
	a string
}

// ParseCSV reads a CSV file and converts it into a slice of problems.
func ParseCSV(file io.Reader) ([]problem, error) {
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse the provided CSV file: %v", err)
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("no problems found in the provided CSV file")
	}
	return parseLines(lines), nil
}

// parseLines processes CSV lines and returns a slice of problem structs.
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		if len(line) < MinColumns {
			log.Fatalf("Invalid CSV format in line %d: each line must have at least two columns.", i+1)
		}
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret

}

// checkAnswer verifies if the provided answer matches the correct answer.
func checkAnswer(a string, q string) bool {
	return strings.TrimSpace(a) == strings.TrimSpace(q)
}

// RunQuiz runs the quiz by asking questions and checking answers against a timer.
func RunQuiz(timer <-chan time.Time, problems []problem) {
	correct := 0
	scanner := bufio.NewScanner(os.Stdin)
	for i, p := range problems { // Iterate through the questions
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			scanner.Scan()
			answerCh <- scanner.Text()
		}() // Capture user input asynchronously
		select {
		case <-timer: // Time's up
			fmt.Println("\nTime's up!")
			fmt.Printf("You answered %d questions correctly and got %d wrong.\n", correct, len(problems)-correct)
			return
		case answer := <-answerCh: // Check the user's answer
			if checkAnswer(answer, p.a) {
				correct++
			}
		}
	}
	fmt.Printf("You answered %d question correctly and got %d wrong.\n", correct, len(problems)-correct)
}
