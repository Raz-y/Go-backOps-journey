package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
	_ "modernc.org/sqlite" // Import the pure Go SQLite driver
)

// pathURL represents a mapping from a short path to a full URL.
// This struct is used to parse YAML / JSON configurations.
type pathURL struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

// MapHandler returns an http.HandlerFunc that attempts to map any requests to their
// corresponding URL based on the provided map. If the path is not found in the map,
// the fallback http.Handler will be called.
func MapHandler(urlPaths map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Extract the request path
		path := req.URL.Path
		// Check if the path exists in the map
		if dest, ok := urlPaths[path]; ok {
			// If the path is found, redirect to the mapped URL
			http.Redirect(res, req, dest, http.StatusFound)
			return
		}
		// If the path is not found, call the fallback handler
		fallback.ServeHTTP(res, req)
	}
}

// SQLiteHandler returns an http.HandlerFunc that attempts to map any requests to their
// corresponding URL based on the provided SQLite database. If the path is not found in the db,
// the fallback http.Handler will be called.
func SQLiteHandler(db *sql.DB, fallback http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Extract the request path
		path := req.URL.Path

		var url string
		// Query the database for the corresponding URL
		// Parameterized query is used to prevent SQL injectiond
		err := db.QueryRow("SELECT url FROM path_urls WHERE path = ?", path).Scan(&url)
		if err == sql.ErrNoRows {
			// If no matching path is found, call the fallback handler
			fallback.ServeHTTP(res, req)
			return
		} else if err != nil {
			// Handle other potential database errors
			log.Printf("Database query error: %v", err)
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Redirect to the URL found in the database
		http.Redirect(res, req, url, http.StatusFound)
	}
}

// YAMLHandler parses the provided YAML configuration and returns an http.HandlerFunc
// that maps any requests to their corresponding URL. If the path is not found in the YAML,
// the fallback http.Handler will be called.
func YAMLHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Handle empty YAML data
	if len(yamlData) == 0 {
		return nil, fmt.Errorf("empty YAML file provided")
	}

	// Parse the YAML data into a slice of pathURL structs.
	parseYaml, err := parseYAML(yamlData)
	if err != nil {
		return nil, err
	}
	// Convert the parsed data into a map.
	pathMap := buildMap(parseYaml)
	// Return the MapHandler with the constructed path map.
	return MapHandler(pathMap, fallback), nil
}

// JSONHandler parses the provided JSON configuration and returns an http.HandlerFunc
// that maps any requests to their corresponding URL. If the path is not found in the JSON,
// the fallback http.Handler will be called.
func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Handle empty JSON data
	if len(jsonData) == 0 {
		return nil, fmt.Errorf("empty JSON file provided")
	}
	// Unmarshal the JSON data into a slice of pathURL structs
	var pathUrls []pathURL
	err := json.Unmarshal(jsonData, &pathUrls)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	// Convert the parsed data into a map
	pathMap := buildMap(pathUrls)
	// Return the MapHandler with the constructed path map
	return MapHandler(pathMap, fallback), nil
}

// parseYAML takes a byte slice of YAML data and unmarshals it into a slice of pathURL structs.
func parseYAML(data []byte) ([]pathURL, error) {
	var pathUrls []pathURL
	// Unmarshal the YAML data into the pathUrls slice.
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}
	return pathUrls, nil
}

// buildMap converts a slice of pathURL structs into a map for quick URL lookup.
func buildMap(data []pathURL) map[string]string {
	mapUrls := make(map[string]string)
	// Iterate over the slice and populate the map.
	for _, pu := range data {
		mapUrls[pu.Path] = pu.URL
	}
	return mapUrls
}
