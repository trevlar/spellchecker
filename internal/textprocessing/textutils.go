package textprocessing

import (
	"strings"
)

/**
 * Checks if the word is possibly a proper noun.
 * @param word The word to check.
 * @return True if the word is possibly a proper noun it returns false otherwise.
 */
func IsPossibleProperNoun(word string) bool {
	// Check if first letter is uppercase
	firstLetter := word[0:1]
	return strings.ToUpper(firstLetter) == firstLetter
}
