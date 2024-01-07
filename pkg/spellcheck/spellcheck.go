package spellcheck

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"spellchecker/config"
	"spellchecker/internal/dictionary"
	"spellchecker/internal/stringutils"
	"strings"
)

func CheckSpellingForFile(filename string, dict *dictionary.Dictionary) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file to check")
		return err
	}
	defer file.Close()

	checker := &SpellChecker{
		dictionary: dict,
		semaphore:  make(chan struct{}, config.MaxGoRoutines),
		spellcheckResults: SpellCheckResults{
			misspellingResults: make(map[string]MisspellingResult),
			misspelledWords:    make([]string, 0),
		},
		hasMisspellings: 			false,
		wordsProcessed:      make(map[string]bool),
	}

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		checker.processLine(scanner.Text(), lineNum)
		lineNum++
	}

	checker.waitGroup.Wait()

	sort.Slice(checker.spellcheckResults.misspelledWords, func(i, j int) bool {
		return checker.spellcheckResults.misspelledWords[i] < checker.spellcheckResults.misspelledWords[j]
	})

	if checker.hasMisspellings {
		printSpellcheckResults(checker.spellcheckResults)
	} else {
		fmt.Println("No misspellings found")
	}
	return nil
}

func findSuggestedWords(word string, dictionary []string) (matchingWords []WordSuggestion) {
	for _, dictWord := range dictionary {
		lowerCaseWord := strings.ToLower(word)

		isCloselyMatchingWord, distance := getDistanceForCloselyMatchingWord(lowerCaseWord, dictWord)

		if isCloselyMatchingWord {
			rank := calculateRank(*distance, lowerCaseWord, dictWord)

			matchingWords = append(matchingWords, WordSuggestion{
				word: dictWord,
				rank: rank,
			})
		}
	}

	sort.Slice(matchingWords, func(i, j int) bool {
		return matchingWords[i].rank < matchingWords[j].rank
	})

	if len(matchingWords) > 10 {
		matchingWords = matchingWords[:10]
	}

	return
}

/**
 * Calculates the rank for a suggested word discovered by damerau levenshtein distance algorithm.
 * The rank goes up for those that have the same starting character as the misspelled word and
 * those that are of the same length.
 */
func calculateRank(distance int, misspelledWord, dictWord string) int {
	rank := distance * 10

	firstLetter := string(misspelledWord[0])
	if strings.HasPrefix(dictWord, firstLetter) {
		rank -= 10
	}
	if len(dictWord) == len(misspelledWord) {
		rank -= 5
	}
	lastLetter := string(misspelledWord[len(misspelledWord)-1])
	if strings.HasSuffix(dictWord, lastLetter) {
		rank -= 5
	}

	return rank
}

func getDistanceForCloselyMatchingWord(word string, dictWord string) (bool, *int) {
	distance := stringutils.DamerauLevenshteinDistance(word, dictWord)

	const similarityThreshold = 3

	if distance <= similarityThreshold {
		return true, &distance
	} else {
		return false, nil
	}
}
