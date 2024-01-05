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
	processing        map[string]bool
	processMutex      sync.Mutex
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
			if len(cleanWord) == 0 {
				continue
			}

			if textprocessing.IsPossibleProperNoun(cleanWord) {
				continue
			}

			if _, matchFound := checker.dictionary.Map[cleanWord]; matchFound {
				continue
			}

			startIndex := max(0, index-3)
			endIndex := min(len(words), index+3+1)

			wordsBefore := strings.Join(words[startIndex:index], " ")
			wordsAfter := strings.Join(words[index+1:endIndex], " ")

			checker.processMutex.Lock()
			if _, isProcessingWord := checker.processing[cleanWord]; !isProcessingWord {
				checker.processing[cleanWord] = true
				checker.processMutex.Unlock()

				checker.waitGroup.Add(1)
				checker.semaphore <- struct{}{}

				go func(cleanWord, wordsBefore, wordsAfter string, lineNum int, index int) {
					defer checker.waitGroup.Done()
					defer func() { <-checker.semaphore }()
					defer func() {
						checker.processMutex.Lock()
						delete(checker.processing, cleanWord)
						checker.processMutex.Unlock()
					}()

					suggestedWords := findSuggestedWords(cleanWord, checker.dictionary.Words)

					checker.mutex.Lock()
					checker.hasMisspellings = true
					checker.spellcheckResults.updateMisspellingResults(cleanWord, wordsBefore, wordsAfter, lineNum, index, suggestedWords)
					checker.mutex.Unlock()
				}(cleanWord, wordsBefore, wordsAfter, lineNum, index)
			} else {
				checker.processMutex.Unlock()
				if existingEntry, exists := checker.spellcheckResults.misspellingResults[cleanWord]; exists {
					updatedEntry := existingEntry
					updatedEntry.wordContext = append(existingEntry.wordContext, WordContext{
						wordsBefore: wordsBefore,
						wordsAfter:  wordsAfter,
						lineNumber:  lineNum,
						wordNumber:  index + 1,
					})
					checker.spellcheckResults.misspellingResults[cleanWord] = updatedEntry
					continue
				}
			}
		}
	}
}
