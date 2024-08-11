package main

import (
	quiz "Quiz/QuizLogic"
	"flag"
	"log"
	"os"
	"time"
)

// main is the entry point of the Quiz application.
func main() {
	// Command-line flags for CSV file path and quiz timer duration.
	csvPtr := flag.String("csv", "Problems.csv", "a csv file in the format of 'question,answer'")
	timerPtr := flag.Int("timer", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// Open the specified CSV file.
	file, err := os.Open(*csvPtr)
	if err != nil {
		log.Fatalf("Failed to open the CSV file: %s", *csvPtr)
	}
	defer file.Close() // Ensure file is closed after use.

	// Parse the CSV file into a list of problems.
	problems, err := quiz.ParseCSV(file)
	if err != nil {
		log.Fatal(err)
	}

	// Create a timer for the quiz duration.
	timer := time.NewTimer(time.Duration(*timerPtr) * time.Second)
	// Start the quiz with the parsed problems and timer.
	quiz.RunQuiz(timer.C, problems)
}
