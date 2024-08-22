# Sitemap Generator
**Sitemap Generator** is a Go package and command-line tool for creating XML sitemaps by crawling websites. It validates URLs, performs website crawls up to a specified depth, and generates sitemaps that comply with the sitemap protocol.


## Table of Contents
- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Command Line](#command-line)
  - [Flags](#flags)
- [Examples](#examples)
- [Contributing](#contributing)
- [Acknowledgements](#acknowledgements)

## Introduction
This project was developed to automate the generation of XML sitemaps by crawling a website. It serves as a practical exercise in Go programming, covering web crawling, URL validation, and XML generation. The project leverages Go's robust standard library to ensure efficiency and reliability.


## Project Structure
```plaintext
SitemapGenerator/
├── htmllink/                 # External package for parsing HTML links (GitHub repository)
│   └── htmllink.go
├── sitemap/                  # Sitemap generation logic
│   ├── sitemap.go            # Core library for generating sitemaps
│   └── sitemap_test.go       # Unit tests for the sitemap library
├── go.mod                    # Module definition file
├── go.sum                    # Module dependencies checksum file
├── main.go                   # Command-line interface for the sitemap generator
├── README.md                 # Project documentation
└── sitemap.xml               # Example output of the generated sitemap (created after running the tool with fitwtech.com)
```
## Features
- **URL Validation**: Ensures that URLs are valid and correctly formatted.
- **Website Crawling**: Recursively crawls websites up to a specified depth to collect URLs.
- **XML Sitemap Generation**: Produces a well-formed XML sitemap that complies with the sitemap protocol.
- **Error Handling**: Comprehensive error handling for robust and reliable sitemap generation.

## Installation
To install and use the Sitemap Generator, follow these steps:

### Clone the repository:

```sh
git clone https://github.com/yourusername/sitemap-generator.git
cd sitemap-generator
```
### Install dependencies:
The project uses Go modules, so ensure Go is installed, then run:

```sh
go mod tidy
```

### Build the project:
You can build the project using:

```sh
go build -o sitemap-generator
```

## Usage
The Sitemap Generator can be used as a command-line tool to crawl a website and generate a sitemap.

### Command Line
To generate a sitemap for a website:

```sh
./sitemap-generator -site=example.com -depth=3
```

### Flags
- **site**: The domain to build the sitemap from. Example: google.com or https://google.com.
- **depth**: The maximum depth to crawl for the sitemap (default is 5).


## Examples
Here are some examples of running the Sitemap Generator:

### Basic usage:

```sh
./sitemap-generator -site=fitwtech.com
```
This command crawls fitwtech.com and generates a sitemap up to the default depth of 5.

### Specifying depth:

```sh
./sitemap-generator -site=fitwtech.com -depth=2
```
This command crawls fitwtech.com up to a depth of 2, collecting and generating the sitemap for URLs found within that depth.

## Contributing
**Contributions are welcome!** If you'd like to contribute to this project, please fork the repository and submit a pull request. Whether it's improving documentation, fixing bugs, or adding new features, your contributions are appreciated.

## Acknowledgements
- ***Inspiration***: This project was inspired by the principles taught in the [Gophercises course](https://gophercises.com/) with a few tweaks of my own.
- ***Dependencies***: The project relies on the `htmllink` package for parsing HTML links, which is available as a separate GitHub folder - [htmllink](https://github.com/Raz-y/Go-backOps-journey/tree/main/HTML_Link_Parser).


