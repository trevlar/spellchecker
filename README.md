# Spellchecker

![Example Image](screenshot/spellchecker.png)

## Overview

This is a simple spellchecker written as a terminal program in Golang. This project demonstrates my ability to write a basic spellchecking application. When looping through the words this program uses goroutines to speed up the process by using multiple threads. There are mutex locks that prevent the same word from being processed at the same time. 

The algorithm chosen for this application uses a hashmap to find exact matches on words. If an exact match is not found then it will use the Damerau-Levenshtein distance algorithm to find matches that are within a certain edit distance. It then will use a calculated rank based on the first and last letter of the word matching and the length of the word to determine the best matches. The application trims the suggested words down to ten.

## Installation

Ensure that you have [Golang](https://golang.org/doc/install) installed on your machine.

To install the spellchecker run the following command in your terminal.

```bash
cd /directory/to/spellchecker
go install ./cmd/spellchecker
```

## Usage

After installing the spellchecker you can run it using the following command with the `<file to check>` and `<dictionary>` arguments. 

```bash
spellchecker <file-to-check> <dictionary>
```

### Arguments

- `<file-to-check>` - The path to the txt file that you want to spell check.
- `<dictionary>` - The path to the dictionary txt file that you want to use to check the spelling of the file.


### Example

```bash
spellchecker ./testdata/incorrect-spelling.txt ./test/dictionary.txt
```

## Testing

This Application includes unit and integration tests that verify the functionality that has been implemented.

To run the tests use the following command.

```bash
go test ./...
```

## Assumptions

- This application is primarily written for the English language using ASCII characters.
- Unicode characters used by other languages will not be properly checked.

## Structure

The application is structured as follows:

```
spellchecker
├── cmd
│   └── spellchecker
│       └── the main entry point for the app (spellchecker.go)
├── config
│   └── central configuration (config.go)
├── internal
│   ├── dictionary
│   │   ├── dictionary processing (dictionary.go)
│   │   └── unit tests for dictionary (dictionary_test.go)
│   ├── stringutils
│   │   ├── utility functions for string manipulation (dalev distance algo) (stringutils.go)
│   │   └── unit tests for string utilities (stringutils_test.go)
│   └── textprocessing
│       ├── text processing and cleanup (textprocessing.go)
│       └── unit tests for text processing (textprocessing_test.go)
├── pkg
│   └── spellcheck
│       ├── core spellchecking logic (spellcheck.go, spellchecker.go)
│       ├── result presentation (result_printer.go)
│       └── unit tests for spellchecking components
├── testdata
│   └── sample files for testing and demonstration
├── dictionary.txt
│   └── sample dictionary file
```

### Key Components

- **cmd/spellchecker**: Contains the main entry point of the application. This is where the spellchecking process is initiated.
- **config**: Stores the global configuration settings which centralizes parameters used while running the spellcheck.
- **internal/dictionary**: Handles opening the dictionary file and returning a map and slice of the dictionary words.
- **internal/stringutils and textprocessing**: Utilities for string manipulation and text processing used to clean and preparing text data for spellchecking. Handles words with dashes, slashes, and apostrophes also built to work with contraction like can't, don't, and won't.
- **pkg/spellcheck**: The spellchecking file is opened at this point. The primary functions for looping through the words is included here. It also includes the components that generate output in the terminal to present the results.
- **testdata**: Contains sample text files for testing the spellchecker.
- **dictionary.txt**: The dictionary file provided for this task.