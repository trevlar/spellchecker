package spellcheck

import (
	"fmt"
	"spellchecker/config"
	"strings"
)

type SpellCheckPrinter struct {
	hasMisspellings bool
	misspelledWords []string
	results map[string]MisspellingResult
}

func (print *SpellCheckPrinter) printResults() {
	if print.hasMisspellings {	
		print.header()
		print.details()
	} else {
		fmt.Println("No misspellings found")
	}
}

func (print *SpellCheckPrinter) header() {
	print.printSeparator("Summary of Misspelled Words:")
	var builder strings.Builder
	for i, word := range print.misspelledWords {
		builder.WriteString(fmt.Sprintf("%s%s%s", config.Red, word, config.Reset))
		if i < len(print.misspelledWords) - 1 {
			builder.WriteString(", ")
		}
	}
	fmt.Println(builder.String())
}

func (print *SpellCheckPrinter) details() {
	print.printSeparator("Spellcheck Details:")
	wordNumber := 1
	for _, word := range print.misspelledWords {
		result := print.results[word]
		fmt.Printf("   %d. \t%sMisspelled Word:%s '%s%s%s'\n", wordNumber, config.Yellow, config.Reset, config.Red, result.word, config.Reset)
		print.occurrences(result.word, result.wordContext)
		print.suggestions(result.word, result.suggestedWords)
		fmt.Println()
		fmt.Println()
		wordNumber++
	}
}

func (print *SpellCheckPrinter) occurrences(word string, wordContext []WordContext) {
	fmt.Printf("\t%sOccurrences:%s\n", config.Yellow, config.Reset)
	currentLetter := 'a'
	for _, wordContext := range wordContext {
		print.printOccurrence(word, currentLetter, wordContext)
		if currentLetter < 'z' {
			currentLetter++
		}
	}
}

func (print * SpellCheckPrinter) printOccurrence(word string, letter rune, context WordContext) {
	fmt.Printf("\t\t%c, %sLine%s %d, %sWord%s %d\n",
		letter,
		config.Yellow,
		config.Reset,
		context.lineNumber,
		config.Yellow,
		config.Reset,
		context.wordNumber,
	)
	fmt.Printf("\t\t   %sContext:%s \"... %s %s%s%s %s\" \n",
		config.Yellow, config.Reset,
		context.wordsBefore,
		config.Yellow, word, config.Reset,
		context.wordsAfter,
	)
}

func (print *SpellCheckPrinter) suggestions(word string, suggestedWords []WordSuggestion) {
	if len(suggestedWords) > 0 {
		var builder strings.Builder
		for i, suggestion := range suggestedWords {
			builder.WriteString(fmt.Sprintf("%s%s%s", config.GreenTxt, suggestion.word, config.Reset))
			if i < len(print.misspelledWords) - 1 {
				builder.WriteString(", ")
			}
		}
		fmt.Printf("\t%sSuggestions:%s %s", config.Yellow, config.Reset, builder.String())
	} else {
		fmt.Printf("\t%sNo suggested words found%s\n", config.Red, config.Reset)
	}
}

func (print *SpellCheckPrinter) printSeparator(title string) {
	fmt.Println()
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println(title)
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println()
}