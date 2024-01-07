package spellcheck

import (
	"fmt"
	"spellchecker/config"
	"strings"
)

type SpellCheckPrinter struct {
	hasMisspellings bool
	results *SpellCheckResults
}

func (print *SpellCheckPrinter) printResults() {
	if print.hasMisspellings {	
		print.header(print.results.misspelledWords)
		print.details()
	} else {
		fmt.Println("No misspellings found")
	}
}

func (print *SpellCheckPrinter) header(misspelledWords []string) {
	fmt.Println("========================================")
	fmt.Println("Summary of Misspelled Words:")
	fmt.Println("========================================")
	fmt.Println()
	concat := fmt.Sprintf("%s, %s", config.Reset, config.Red)
	fmt.Printf("%s%s%s\n", config.Red, strings.Join(misspelledWords, concat), config.Reset)
	fmt.Println()
	fmt.Println()
}

func (print *SpellCheckPrinter) details() {
	spellcheckResults := print.results

	fmt.Println("========================================")
	fmt.Println("Spellcheck Details:")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println()

	wordNumber := 1
	for _, word := range spellcheckResults.misspelledWords {
		result := spellcheckResults.misspellingResults[word]
		fmt.Printf("   %d. \t%sMisspelled Word:%s '%s%s%s'\n", wordNumber, config.Yellow, config.Reset, config.Red, result.word, config.Reset)
		print.occurrences(result.word, result.wordContext)
		print.suggestions(result.word, result.suggestedWords)
		
		fmt.Println()
		wordNumber++
	}
}

func (print *SpellCheckPrinter) occurrences(word string, wordContext []WordContext) {
	fmt.Printf("\t%sOccurrences:%s\n", config.Yellow, config.Reset)
	currentLetter := 'a'
	for _, wordContext := range wordContext {
		fmt.Printf("\t\t%c, %sLine%s %d, %sWord%s %d\n",
			currentLetter,
			config.Yellow,
			config.Reset,
			wordContext.lineNumber,
			config.Yellow,
			config.Reset,
			wordContext.wordNumber,
		)
		fmt.Printf("\t\t   %sContext:%s \"... %s %s%s%s %s\" \n",
			config.Yellow, config.Reset,
			wordContext.wordsBefore,
			config.Yellow, word, config.Reset,
			wordContext.wordsAfter,
		)
		if currentLetter < 'z' {
			currentLetter++
		}
	}
}

func (print *SpellCheckPrinter) suggestions(word string, suggestedWords []WordSuggestion) {
	if len(suggestedWords) > 0 {
		suggestedWordsList := []string{}
		for _, wordSuggestion := range suggestedWords {
			suggestedWordsList = append(suggestedWordsList, wordSuggestion.word)
		}

		concat := fmt.Sprintf("%s, %s", config.Reset, config.GreenTxt)
		fmt.Printf("\t%sSuggestions:%s %s%s%s", config.Yellow, config.Reset, config.GreenTxt, strings.Join(suggestedWordsList, concat), config.Reset)
		fmt.Println()
	} else {
		fmt.Printf("\t%sNo suggested words found%s\n", config.Red, config.Reset)
	}
}
