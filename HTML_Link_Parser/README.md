# HTMLLink Parser

HTMLLink Parser is a simple Go package that extracts hyperlinks (`<a href="...">`) from HTML documents. This project was a practical exercise in Go programming and a refresher on algorithms, graph theory, and tree data structures.

## Table of Contents

- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Lessons Learned](#lessons-learned)
- [Contributing](#contributing)
- [Acknowledgements](#acknowledgements)

## Introduction

This project was undertaken as a personal challenge to revisit core concepts in algorithms, graph theory, and trees, while also deepening my understanding of API design in Go. By first building the HTMLLink Parser from scratch and then following [Gophercises](https://gophercises.com/) suggestions and building on top of it, I aimed to solidify my knowledge and apply it to a practical and useful tool.

## Project Structure
```plaintext
HTMLLinkParser/
├── examples/
│   ├── index.html            # Sample HTML file for parsing
│   └── main.go               # Example usage of the htmllink package
├── htmllink/
│   ├── htmllink.go           # Core library for parsing HTML links
│   └── htmllink_test.go      # Unit tests for the htmllink library
├── go.mod                    # Module definition file
├── go.sum                    # Module dependencies checksum file
└── README.md                 # Project documentation
```

## Features

- **HTML Link Parsing:** Extracts all hyperlinks from an HTML document.
- **Error Handling:** Robust error handling for invalid or malformed HTML.
- **Flexible Input:** Supports parsing from both file input and raw HTML strings.

## Installation

To install and use the HTMLLink Parser, follow these steps:

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/htmllink-parser.git
    cd htmllink-parser
    ```

2. **Install dependencies:**
    The project relies on Go modules, so ensure you have Go installed, and run:
    ```sh
    go mod tidy
    ```

3. **Build the project:**
    You can build the project using:
    ```sh
    go build -o htmllink-parser
    ```

## Usage 

The HTMLLink Parser can be used directly from the command line or as a Go package.

### Command Line

To parse an HTML file and extract links:

```sh
./htmllink-parser -file=index.html
```

### As a Go Package

To use the htmllink package in your Go projects:

```go
package main

import (
    "fmt"
    "log"
    "os"
    htmllink "htmllink/htmllink"
)

func main() {
    file, _ := os.Open("index.html")
    defer file.Close()

    links, err := htmllink.ParseLinks(file)
    if err != nil {
        log.Fatalf("Failed to parse HTML: %v", err)
    }
    fmt.Printf("%+v\n", links)
}

```

## Examples
Here are some examples of what the output might look like when using the parser:

- Single Link:
```html
<a href="/home">Home</a>
```
- Output:
``` plaintext
[{Href:/home Text:Home}]
```

- Multiple Links:
```html
<a href="/about">About</a> <a href="/contact">Contact</a>
```
Output:
``` plaintext
[{Href:/about Text:About} {Href:/contact Text:Contact}]
```

- Nested Links:
```html
<a href="/outer">Outer <a href="/inner">Inner</a></a>
```
Output:
``` plaintext
[{Href:/outer Text:Outer} {Href:/inner Text:Inner}]
```

## Lessons Learned

Throughout this project, I gained valuable insights into several key areas of software development:

### Algorithms and Data Structures

- Revisiting graph theory and tree structures reminded me of the importance of understanding the underlying data structures when working with HTML parsing. Each node in the DOM can be seen as part of a tree, making it essential to navigate and manipulate these nodes effectively.
  
### Debugging Techniques

- **First Debugging Lesson:** The initial struggles with parsing and edge cases underscored the importance of thorough debugging and the need to anticipate potential issues in complex HTML structures.

### API Design
- **Learning about API Design:** This project reinforced the importance of designing intuitive and user-friendly APIs. I considered what methods and functions users would likely need and structured the API to be as accessible and logical as possible. The focus was on simplicity, clear error messages, and flexibility in input types.
  
## Contributing

Contributions are welcome! If you would like to contribute to the development of this project, please fork the repository and submit a pull request. Whether it's improving documentation, fixing bugs, or adding new features, all contributions are appreciated.

## Acknowledgements
- ***Inspiration:*** This project was inspired by the Go course on [Gophercises](https://gophercises.com/).
- ***Dependencies:*** This project relies on the `golang.org/x/net/html` package for HTML parsing.
