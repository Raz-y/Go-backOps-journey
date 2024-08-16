package main

import (
	"flag"
	"fmt"
	htmllink "htmllink/htmllink"
	"log"
	"os"
)

func main() {
	// Command-line flag for the input HTML file.
	file := flag.String("file", "index.html", "Specify the HTML file to be parsed")
	flag.Parse()

	if *file == "" {
		log.Fatal("Error: No file provided. Please use the -file flag to specify an HTML file.")
	}

	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("Error: Could not open file '%s'. Error: %v", *file, err)
	}
	defer f.Close() // Ensure the file is closed when the function exits.

	// Parse the HTML file to extract hyperlinks.
	links, err := htmllink.ParseLinks(f)
	if err != nil {
		log.Fatalf("Error: Failed to parse HTML from file '%s'. Ensure the file contains valid HTML content. Detailed error: %v", *file, err)
	}

	// Output the extracted links in a formatted manner.
	fmt.Printf("Extracted links from '%s':\n%+v\n", *file, links)
}
