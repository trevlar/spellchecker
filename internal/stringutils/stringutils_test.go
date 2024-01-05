package stringutils

import (
	"fmt"
	"testing"
)

func TestDamerauLevenshteinDistance(t *testing.T) {
	testCases := []struct {
		name             string
		sourceString     string
		targetString     string
		expectedDistance int
	}{
		{"Four insertion", "hello", "h", 4},
		{"Three insertions", "hello", "he", 3},
		{"Two insertions", "hello", "hel", 2},
		{"One insertions", "hello", "hell", 1},
		{"Four deletion", "h", "hello", 4},
		{"Three deletions", "he", "hello", 3},
		{"Two deletions", "hel", "hello", 2},
		{"One deletions", "hell", "hello", 1},
		{"One substitution", "hello", "hallo", 1},
		{"Two substitutions", "hello", "hailo", 2},
		{"Three substitutions", "hello", "hailu", 3},
		{"Four substitutions", "hello", "haiku", 4},
		{"Two transpositions", "haiuk", "haiku", 2},
		{"Three transpositions", "hiauk", "haiku", 3},
		{"Four transpositions", "kiauh", "haiku", 4},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("should return %d when comparing strings with %s", testCase.expectedDistance, testCase.name)

		t.Run(testName, func(t *testing.T) {
			distance := DamerauLevenshteinDistance(testCase.sourceString, testCase.targetString)
			if distance != testCase.expectedDistance {
				t.Errorf("Expected distance to be %d, got %d", testCase.expectedDistance, distance)
			}
		})
	}

	t.Run("should return 0 when comparing empty strings", func(t *testing.T) {
		sourceString := ""
		targetString := ""

		distance := DamerauLevenshteinDistance(sourceString, targetString)

		if distance != 0 {
			t.Errorf("Expected distance to be 0, got %d", distance)
		}
	})

	t.Run("should return 0 when comparing identical strings", func(t *testing.T) {
		sourceString := "hello"
		targetString := "hello"

		distance := DamerauLevenshteinDistance(sourceString, targetString)

		if distance != 0 {
			t.Errorf("Expected distance to be 0, got %d", distance)
		}
	})
}
