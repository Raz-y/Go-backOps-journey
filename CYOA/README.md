# Choose Your Own Adventure (CYOA) Web & CLI Application

This project is a Go-based implementation of a "Choose Your Own Adventure" (CYOA) story application. It allows users to experience a CYOA story either through a web interface or directly in the command line.

## Project Structure

The project is organized into the following directories and files:

``` plaintext
CYOA/
|-- cyoa/
|   |-- cyoa.go         # Core logic for the CYOA story handling
|   |-- cli.go          # CLI logic for navigating the story in the command line
|   |-- handler.go      # HTTP handler logic for serving the story via a web server
|   |-- templates.go    # Template loading and management for rendering HTML
|-- templates/
|   |-- story.html      # Main HTML template for rendering the story
|   |-- default.html    # Default HTML template for fallback
|-- main.go             # Entry point of the application
|-- story.json          # Example story in JSON format
```
### 1. `main.go`

This file serves as the entry point for the application. Depending on the command-line flags, it will either start an HTTP server to serve the story as a web application or run the story in the command-line interface (CLI).

### 2. cyoa/ Directory

This directory contains the core logic of the application:

- `cyoa.go`: Defines the core structures (`Story`, `Chapter`, `Option`) and provides functions for loading a story from a JSON file.
- `cli.go`: Implements the CLI flow for navigating the story directly in the terminal.
- `handler.go`: Implements the HTTP handler logic, allowing the story to be served via a web server.
- `templates.go`: Manages the loading and parsing of HTML templates used in the web interface.

### 3. templates/ Directory

This directory contains the HTML templates used for rendering the story in the web interface:

- `story.html`: The main template used for displaying the story chapters and options.
- `default.html`: A fallback template in case the main template is not available.

### 4. story.json

An example JSON file that defines a CYOA story. This file is loaded by the application and defines the structure of the story, including chapters and options.

## Usage

### 1. Running the Web Server
To run the application as a web server, use the following command:
```bash
go run main.go -file=story.json
```
This will start an HTTP server on the specified port (default is `3000`) and serve the story as a web application. You can access it by navigating to `http://localhost:3000/story/` in your browser.

### 2. Running the CLI

To run the application in CLI mode, use the following command:
```bash
go run main.go -cli -file=story.json
```
This will run the story directly in your terminal, allowing you to navigate through the chapters by selecting options.

### 3. Command-Line Flags

- `-port`: Specifies the port on which the web server will run. Default is `3000`.
- `-cli`: Enables CLI mode. If this flag is set, the story will run in the terminal instead of being served via HTTP.
- `-file`: Specifies the path to the JSON file containing the story.

### 4. Example Story JSON
Every story should start with an `"intro"` chapter, as the application is designed to begin the narrative from this point. The `"intro"` chapter acts as the entry point to the story.

Here's an example of how a CYOA story might be structured in the story.json file:

```json
{
  "intro": {
    "title": "The Beginning",
    "story": ["It was a dark and stormy night...", "You see a light in the distance."],
    "options": [
      {"text": "Walk towards the light", "arc": "light"},
      {"text": "Stay where you are", "arc": "darkness"}
    ]
  },
  "light": {
    "title": "The Light",
    "story": ["The light grows brighter...", "You find yourself at the entrance of a cave."],
    "options": [
      {"text": "Enter the cave", "arc": "cave"},
      {"text": "Run away", "arc": "escape"}
    ]
  }
  // Add more arcs as needed...
}
```
## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to submit a pull request or open an issue.

To set up the project locally:

1. Clone the repository.
2. Run the application using `go run main.go -file=story.json`(add `-cli` for CLI mode).

## Acknowledgements

- Inspired by the Go course on [Gophercises](https://gophercises.com/).
