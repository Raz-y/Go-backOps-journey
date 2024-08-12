# URL Shortener 🔗

A simple and efficient URL shortener service written in Go. This project allows you to map short paths to full URLs, either through a static map, configuration files (YAML/JSON), or an SQLite database.

## Project Structure 📂
```plaintext
URL_Shortner/
├── go.mod                    # Module definition file
├── go.sum                    # Module dependencies checksum file
├── handler/
│   ├── handlers.go           # Handlers for URL mapping and configuration
│   └── handlers_test.go      # Unit tests for the handlers
├── main.go                   # Entry point for the application
├── pathToUrlMapping.db       # Example SQLite database file for URL mappings
├── pathToUrlMapping.json     # Example JSON file for URL mappings
└── pathToUrlMapping.yaml     # Example YAML file for URL mappings
```

## Features ✨

- **Multiple Configuration Options**: Use a static map, YAML, JSON, or an SQLite database to configure your URL mappings.
- **Fallback Support**: If a path isn't found in your configuration, the server will fall back to a default handler.
- **Extensible Handlers**: Easy to extend with additional handlers for different data sources.
  
## Getting Started 🚀

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/URL_Shortner.git
   cd URL_Shortner
   ```
   
2. **Install dependencies:**

   The project uses Go modules, so dependencies are managed automatically. Simply run: 
   ```bash
   go mod tidy
   ```

## Configuration
You can configure the URL Shortener using one of the following methods:

- **Static Map:** Modify the pathsToUrl map in main.go for hardcoded mappings.
- **YAML or JSON File:** Provide a YAML or JSON file with your path-to-URL mappings.
- **SQLite Database:** Use a SQLite database file to manage your URL mappings dynamically.

> Note: Only one configuration method should be used at a time.


## Running the Application
To run the server, you can specify a configuration file or a SQLite database using command-line flags:
- **Using a YAML or JSON File:**
  
  ```bash
  go run main.go --config=pathToUrlMapping.yaml
  ```
- **Using a SQLite Database:**
  
  ```bash
  go run main.go --db=pathToUrlMapping.db
  ```
- **Without any configuration file (default map):**
  
  ```bash
  go run main.go
  ```
The server will start on port 3000.

## Example Usage
Once the server is running, you can shorten URLs by navigating to the short paths:

- Navigate to http://localhost:3000/raz-go and you'll be redirected to https://github.com/Raz-y/Go-backOps-journey.
- Navigate to http://localhost:3000/non-existent to trigger the fallback handler.

To add a new mapping:
- **Static Map**: Modify `pathsToUrl` in `main.go` by adding a new key-value pair.
- **YAML File**: Add a new entry in the `pathToUrlMapping.yaml`.
- **JSON File**: Add a new entry in the `pathToUrlMapping.json`.
- **SQLite Database**: Insert a new record into the `path_urls` table in the database.


## Running Tests 
To run the unit tests, execute:
```bash
go test ./handler/...
```

## Contributing 🤝
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Acknowledgements
- Inspired by the Go course on [Gophercises](https://gophercises.com/).
- [modernc.org/sqlite](modernc.org/sqlite) - for the pure Go SQLite driver.
