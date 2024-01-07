package spellcheck

import (
	"spellchecker/internal/dictionary"
	"spellchecker/internal/textprocessing"
	"strings"
	"sync"
)

type SpellChecker struct {
	dictionary        *dictionary.Dictionary
	spellcheckResults SpellCheckResults
	wordsProcessed    map[string]bool
	wordsProcMutex		sync.Mutex
	waitGroup         sync.WaitGroup
	mutex             sync.Mutex
	semaphore         chan struct{}
	hasMisspellings   bool
}

func (checker *SpellChecker) processLine(line string, lineNum int) {
	words := strings.Fields(line)
	for index, word := range words {

		cleanWords := textprocessing.CleanWord(word)
		for _, cleanWord := range cleanWords {

			if len(cleanWord) == 0 || textprocessing.IsPossibleProperNoun(cleanWord) {
				continue
			}

			if _, matchFound := checker.dictionary.Map[cleanWord]; !matchFound {
				checker.waitGroup.Add(1)
				checker.semaphore <- struct{}{}

				checker.hasMisspellings = true
				go checker.handleMisspelling(cleanWord, words, index, lineNum)
			}
		}
	}
}

func (checker *SpellChecker) handleMisspelling(cleanWord string, words []string, index, lineNum int) {
	defer checker.waitGroup.Done()
	defer func() { <-checker.semaphore }()

	startIndex := max(0, index-3)
	endIndex := min(len(words), index+3+1)

	wordsBefore := strings.Join(words[startIndex:index], " ")
	wordsAfter := strings.Join(words[index+1:endIndex], " ")

	checker.wordsProcMutex.Lock()
	if _, isProcessingWord := checker.wordsProcessed[cleanWord]; !isProcessingWord {
		checker.wordsProcessed[cleanWord] = true
		checker.wordsProcMutex.Unlock()

		suggestedWords := findSuggestedWords(cleanWord, checker.dictionary.Words)
		
		checker.mutex.Lock()
		checker.spellcheckResults.updateMisspellingResults(cleanWord, wordsBefore, wordsAfter, lineNum, index, suggestedWords)
		checker.mutex.Unlock()
	} else {
		if existingEntry, exists := checker.spellcheckResults.misspellingResults[cleanWord]; exists {
			updatedEntry := existingEntry
			updatedEntry.wordContext = append(existingEntry.wordContext, WordContext{
				wordsBefore: wordsBefore,
				wordsAfter:  wordsAfter,
				lineNumber:  lineNum,
				wordNumber:  index + 1,
			})
			checker.spellcheckResults.misspellingResults[cleanWord] = updatedEntry
		}
		checker.wordsProcMutex.Unlock()
	}
}
