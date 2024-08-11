# Quiz Application

This is a simple command-line quiz application written in Go. The application reads quiz questions and answers from a CSV file and presents them to the user in a timed quiz format. The user has a limited amount of time to answer each question, and the application tracks the number of correct and incorrect answers.

## Features

- Load quiz questions from a CSV file.
- Set a custom time limit for the quiz.
- Track the number of correct and incorrect answers.
- Display results at the end of the quiz.

## Project Structure
```plaintext
Quiz/
├── QuizLogic/
│   ├── quiz.go
│   └── quiz_test.go
├── main.go
└── go.mod
```

## File Descriptions

- **QuizLogic/quiz.go**: Contains the core quiz logic, including CSV parsing and quiz execution.
- **QuizLogic/quiz_test.go**: Contains unit tests for the quiz logic.
- **main.go**: The entry point for the application.
- **go.mod**: Go module file.

## Requirements

- Go 1.21.11 or higher

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/quiz-app.git
    ```
2. Change to the project directory:
    ```bash
    cd quiz-app
    ```
3. Build the application:
    ```bash
    go build
    ```

## Usage

Run the application with default settings:
```bash
./quiz-app
```

Specify a custom CSV file and timer:
```bash
./quiz-app -csv=questions.csv -timer=60
```

## CSV File Format
The CSV file should contain questions and answers in the following format:
```plaintext
question1,answer1
question2,answer2
...
```
Example:
```plaintext
What is 2+2?,4
What is the capital of France?,Paris
```
## Running Tests
To run the unit tests for the quiz logic, use the following command:
```bash
go test ./QuizLogic/...
```
## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Acknowledgements
Inspired by the Go course on [Gophercises](https://gophercises.com/).
