package cyoa

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CLIFlow
func CLIFlow(s Story, c string) {
	scanner := bufio.NewScanner(os.Stdin)

	// Print the title
	fmt.Printf("\n======= %v =======\n\n", s[c].Title)

	// Print the story content
	fmt.Printf("%v\n", strings.Join(s[c].Story, "\n    "))
	fmt.Print("\n-----------------------------\n")

	// If there are no options, end the story
	if len(s[c].Options) == 0 {
		fmt.Println("The End. Thanks for playing!")
		return
	}

	// Print the options
	fmt.Println("What would you like to do next?")
	for i, opt := range s[c].Options {
		fmt.Printf("  %d) %s\n", i+1, opt.Text)
	}

	// Add a line break before the input prompt
	fmt.Print("\nChoose an option (or 'q' to quit): ")

	// Get user input
	if scanner.Scan() {
		co := scanner.Text()

		// Allow the user to quit by typing "q"
		if strings.ToLower(co) == "q" {
			fmt.Println("Exiting the story. Thanks for playing!")
			return
		}

		// Convert user input to integer
		choice, err := strconv.Atoi(co)
		if err != nil || choice < 1 || choice > len(s[c].Options) {
			fmt.Println("Invalid choice. Please enter a valid number.")
			CLIFlow(s, c) // Restart the current chapter
			return
		}

		// Recursively call CLIFlow with the next chapter
		nextChapter := s[c].Options[choice-1].NextArc
		CLIFlow(s, nextChapter)
	}
}
