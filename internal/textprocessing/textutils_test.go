package textprocessing

import (
	"testing"
)

func TestIsPossibleProperNoun(t *testing.T) {
	t.Run("should return true if the first letter is uppercase", func(t *testing.T) {
		input := "Test"

		if !IsPossibleProperNoun(input) {
			t.Errorf("Expected true but got false")
		}
	})

	t.Run("should return false if the first letter is lowercase", func(t *testing.T) {
		input := "test"

		if IsPossibleProperNoun(input) {
			t.Errorf("Expected false but got true")
		}
	})
}
