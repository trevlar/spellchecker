package textprocessing

import (
	"strings"
)

func IsPossibleProperNoun(word string) bool {
	// Check if first letter is uppercase
	firstLetter := word[0:1]
	return strings.ToUpper(firstLetter) == firstLetter
}
