package textprocessing

import (
	"regexp"
	"strings"
)

func CleanWord(word string) []string {
	cleanWord := normalizeApostrophe(word)
	cleanWord = removePossessive(cleanWord)
	cleanWord = removePunctuation(cleanWord)
	if strings.Contains(cleanWord, "'") {
		cleanWords := expandContraction(cleanWord)
		return strings.Fields(cleanWords)
	}
	if strings.Contains(cleanWord, "-") {
		cleanWords := splitByCharacter(cleanWord, "-")
		return cleanWords
	}
	if strings.Contains(cleanWord, "/") {
		cleanWords := splitByCharacter(cleanWord, "/")
		return cleanWords
	}
	return []string{cleanWord}
}

func expandContraction(word string) string {
	contractions := map[string]string{
		"ain't":     "am not",
		"aren't":    "are not",
		"can't":     "cannot",
		"couldn't":  "could not",
		"didn't":    "did not",
		"doesn't":   "does not",
		"don't":     "do not",
		"hadn't":    "had not",
		"hasn't":    "has not",
		"haven't":   "have not",
		"he'd":      "he would",
		"he'll":     "he will",
		"he's":      "he is",
		"i'd":       "I would",
		"i'll":      "I will",
		"i'm":       "I am",
		"i've":      "I have",
		"isn't":     "is not",
		"it's":      "it is",
		"mightn't":  "might not",
		"mustn't":   "must not",
		"she'd":     "she would",
		"she'll":    "she will",
		"she's":     "she is",
		"shouldn't": "should not",
		"they'd":    "they would",
		"they'll":   "they will",
		"they're":   "they are",
		"they've":   "they have",
		"wasn't":    "was not",
		"we'd":      "we would",
		"we'll":     "we will",
		"we're":     "we are",
		"we've":     "we have",
		"weren't":   "were not",
		"won't":     "will not",
		"wouldn't":  "would not",
		"you'd":     "you would",
		"you'll":    "you will",
		"you're":    "you are",
		"you've":    "you have",
	}
	apostropheRegex.ReplaceAllString(word, "'")

	if expandedWord, exists := contractions[word]; exists {
		return expandedWord
	}

	return word
}

var (
	apostropheRegex  = regexp.MustCompile(`’`)
	possessiveRegex  = regexp.MustCompile(`'s$`)
	punctuationRegex = regexp.MustCompile(`^\P{L}+|\P{L}+$`)
)

func splitByCharacter(word string, character string) []string {
	if strings.Contains(word, character) {
		return strings.Split(word, character)
	}
	return []string{word}
}

func normalizeApostrophe(word string) string {
	return apostropheRegex.ReplaceAllString(word, "'")
}

func removePossessive(word string) string {
	return possessiveRegex.ReplaceAllString(word, "")
}

func removePunctuation(word string) string {
	return punctuationRegex.ReplaceAllString(word, "")
}
