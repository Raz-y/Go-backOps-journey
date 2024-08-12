package main

import (
	handlers "URL_Shortner/handler"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite" // Import the pure Go SQLite driver
)

// main is the entry point of the URL Shortner service. It sets up the HTTP server
// and handlers, and starts the server on port 3000.
func main() {
	// Command-line flag for the configuration file path (YAML or JSON)
	configFile := flag.String("config", "", "Path to a configuration file (YAML or JSON) containing path-to-URL mappings. Example: -config=config.yaml")
	dbFile := flag.String("db", "", "Path to a SQLite database file containing path-to-URL mappings. Example: -db=urls.db")
	flag.Parse()

	// Set default mux
	mux := defaultMux()

	pathsToURLs := map[string]string{
		"/raz-backend": "https://github.com/Raz-y/backend-development-roadmap",
		"/raz-go":      "https://github.com/Raz-y/Go-backOps-journey",
	}
	// MapHandler handles the redirection of short URLs
	mapHandler := handlers.MapHandler(pathsToURLs, mux)

	// Determine the file type by extention
	var handler http.HandlerFunc

	if *dbFile != "" {
		db, err := sql.Open("sqlite", *dbFile)
		if err != nil {
			log.Fatalf("Couldn't open SQLite database file %s: %v", *dbFile, err)
		}
		defer db.Close()

		handler = handlers.SQLiteHandler(db, mapHandler)
		if err != nil {
			panic(err)
		}

	} else if *configFile != "" { // Check for given config file
		ext := filepath.Ext(*configFile)
		switch ext {
		case ".yaml", ".yml":
			// Handle YAML file
			// Read the YAML file specified by the user.
			yamlData, err := os.ReadFile(*configFile)
			if err != nil {
				log.Fatalf("Couldn't read YAML file %s: %v", *configFile, err)
			}
			// Parse the YAML data and generate the HTTP handler.
			handler, err = handlers.YAMLHandler(yamlData, mapHandler)
			if err != nil {
				panic(err)
			}
		case ".json":
			// Handle JSON file
			// Read the JSON file specified by the user.
			jsonData, err := os.ReadFile(*configFile)
			if err != nil {
				log.Fatalf("Couldn't read JSON file %s: %v", *configFile, err)
			}
			// Parse the JSON data and generate the HTTP handler.s
			handler, err = handlers.JSONHandler(jsonData, mapHandler)
			if err != nil {
				panic(err)
			}
		default:
			// If the file extension is not recognized, log an error and terminate the program.
			log.Fatalf("Unsupported file extension: %s. Please provide a .yaml, .yml, or .json file.", ext)
		}
	} else { // Check if a file was given
		log.Fatal("No configuration file provided. Use --config to specify a file.")

	}

	fmt.Println("Starting the server on :3000")
	// Start the HTTP server on port 3000 with the YAML handler.
	err := http.ListenAndServe(":3000", handler)
	// Log an error if the server fails to start.
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// defaultMux returns an HTTP request multiplexer with a default "Hello World" handler.
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

// hello is the default handler that responds with "Hello World".
func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World\n")
}
