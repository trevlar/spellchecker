package spellcheck

type WordContext struct {
	wordsBefore string
	wordsAfter  string
	lineNumber  int
	wordNumber  int
}

type WordSuggestion struct {
	word string
	rank int
}

type MisspellingResult struct {
	word           string
	suggestedWords []WordSuggestion
	wordContext    []WordContext
}

type SpellCheckResults struct {
	misspellingResults map[string]MisspellingResult
	misspelledWords    []string
}

func (sc *SpellCheckResults) updateMisspellingResults(cleanWord, wordsBefore, wordsAfter string, lineNum, index int, suggestedWords []WordSuggestion) {
	if existingEntry, exists := sc.misspellingResults[cleanWord]; exists {
		updatedEntry := existingEntry
		updatedEntry.wordContext = append(updatedEntry.wordContext, WordContext{
			wordsBefore: wordsBefore,
			wordsAfter:  wordsAfter,
			lineNumber:  lineNum,
			wordNumber:  index + 1,
		})
		sc.misspellingResults[cleanWord] = updatedEntry
	} else {
		sc.misspelledWords = append(sc.misspelledWords, cleanWord)
		sc.misspellingResults[cleanWord] = MisspellingResult{
			word:           cleanWord,
			suggestedWords: suggestedWords,
			wordContext: []WordContext{
				{
					wordsBefore: wordsBefore,
					wordsAfter:  wordsAfter,
					lineNumber:  lineNum,
					wordNumber:  index + 1,
				},
			},
		}
	}
}
