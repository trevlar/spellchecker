package spellcheck

import (
	"fmt"
	"spellchecker/config"
	"strings"
)

func printSpellcheckResults(spellcheckResults SpellCheckResults) {
	printSpellcheckResultsHeader(spellcheckResults.misspelledWords)
	printSpellcheckResultDetails(spellcheckResults)
}

func printSpellcheckResultsHeader(misspelledWords []string) {
	fmt.Println("========================================")
	fmt.Println("Summary of Misspelled Words:")
	fmt.Println("========================================")
	fmt.Println()
	concat := fmt.Sprintf("%s, %s", config.Reset, config.Red)
	fmt.Printf("%s%s%s\n", config.Red, strings.Join(misspelledWords, concat), config.Reset)
	fmt.Println()
	fmt.Println()
}

func printSpellcheckResultDetails(spellcheckResults SpellCheckResults) {
	fmt.Println("========================================")
	fmt.Println("Spellcheck Details:")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println()

	wordNumber := 1
	for _, word := range spellcheckResults.misspelledWords {
		result := spellcheckResults.misspellingResults[word]
		currentLetter := 'a'
		fmt.Printf("   %d. \t%sMisspelled Word:%s '%s%s%s'\n", wordNumber, config.Yellow, config.Reset, config.Red, result.word, config.Reset)
		fmt.Printf("\t%sOccurrences:%s\n", config.Yellow, config.Reset)
		for _, wordContext := range result.wordContext {
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
				config.Yellow, result.word, config.Reset,
				wordContext.wordsAfter,
			)
			if currentLetter < 'z' {
				currentLetter++
			}
		}

		if len(result.suggestedWords) > 0 {

			suggestedWords := []string{}
			for _, wordSuggestion := range result.suggestedWords {
				suggestedWords = append(suggestedWords, wordSuggestion.word)
			}

			concat := fmt.Sprintf("%s, %s", config.Reset, config.GreenTxt)
			fmt.Printf("\t%sSuggestions:%s %s%s%s", config.Yellow, config.Reset, config.GreenTxt, strings.Join(suggestedWords, concat), config.Reset)
			fmt.Println()
		} else {
			fmt.Printf("\t%sNo suggested words found%s\n", config.Red, config.Reset)
		}
		fmt.Println()
		wordNumber++
	}
}
