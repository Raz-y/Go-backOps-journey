package main

import (
	"CYOA/cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// main.go serves as the entry point for the CYOA web application.
// It reads a JSON story file, initializes the handler with templates, and starts an HTTP server.
func main() {
	// Parse command-line flags
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	cliFlow := flag.Bool("cli", false, "The status if for cli flow")
	storyFile := flag.String("file", "", "Path to a JSON file containing the choose your own adventure story. Example: -config=story.json")
	flag.Parse()

	if *storyFile == "" {
		log.Fatal("No file provided. Use --file to specify a file.")
	}

	// Open and parse the story file
	f, err := os.Open(*storyFile)
	if err != nil {
		log.Fatalf("Error in opening file %v, Error: %v", *storyFile, err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		log.Fatal(err)
	}

	// Check if cli flag exists to initiate cli flow
	if *cliFlow {
		cyoa.CLIFlow(story, "intro") // Start the story from the "intro" chapter
		return
	}

	// Load templates
	tpl := cyoa.MustLoadTemplates("templates")

	// Create the HTTP handler with the custom template and custom path
	h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl), cyoa.WithPathFunc(customPathFn))

	// Register the handler and start the server

	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

// customPathFn is a custom function that defaults to "start" instead of "intro".
func customPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/start" // Custom default chapter
	}
	return path[len("/story/"):] // Remove leading "/story/"
}
