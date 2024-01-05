package main

import (
	"fmt"
	"os"
	"spellchecker/internal/dictionary"
	"spellchecker/pkg/spellcheck"
)

const (
	ExNoInput = 66
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: spellchecker <dictionary.txt> <file-to-check.txt>")
		os.Exit(1)
	}

	dictFileName := os.Args[1]
	fileNameToCheck := os.Args[2]

	dict, err := dictionary.ReadFile(dictFileName)
	if err != nil {
		fmt.Printf("Failed to read dictionary file %s: %v", dictFileName, err)
		os.Exit(ExNoInput)
	}

	if err := spellcheck.CheckSpellingForFile(fileNameToCheck, dict); err != nil {
		fmt.Printf("Error checking spelling in file %s: %v", fileNameToCheck, err)
		os.Exit(ExNoInput)
	}
}
