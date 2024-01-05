package textprocessing

import (
	"testing"
)

func TestCleanWord(t *testing.T) {
	t.Run("should remove a period at the end of a word", func(t *testing.T) {
		input := "testing."
		result := CleanWord(input)

		expected := []string{"test"}
		if len(result) != len(expected) {
			t.Errorf("Expected %d but got %d", len(expected), len(result))
		}
	})

	t.Run("should expand a contraction", func(t *testing.T) {
		input := "can't"
		result := expandContraction(input)

		expected := "cannot"
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	t.Run("should replace curly apostrophes (’) in the word", func(t *testing.T) {
		input := "test’s"
		result := normalizeApostrophe(input)

		expected := "test's"
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	t.Run("should return a clean list of words", func(t *testing.T) {
		input := "test's"
		result := removePossessive(input)

		expected := "test"
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	t.Run("should remove a comma at the end of a word", func(t *testing.T) {
		input := "testing,"
		result := removePunctuation(input)

		expected := "testing"
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

	t.Run("should split words separated by a dash", func(t *testing.T) {
		input := "test-word"
		result := splitByCharacter(input, "-")

		expected := []string{"test", "word"}
		if len(result) != len(expected) {
			t.Errorf("Expected %d but got %d", len(expected), len(result))
		}
	})
}
