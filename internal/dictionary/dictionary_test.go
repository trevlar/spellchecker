package dictionary

import "testing"

func TestReadFile(t *testing.T) {
	t.Run("should return an error if the file does not exist", func(t *testing.T) {
		_, err := ReadFile("../../testdata/does-not-exist.txt")
		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})

	t.Run("should return a slice of strings", func(t *testing.T) {
		dictionary, err := ReadFile("../../testdata/dictionary.txt")
		if err != nil {
			t.Error("Unexpected error occurred while reading file")
		}

		if len(dictionary.Words) == 0 {
			t.Error("expected words to be returned but got none")
		}

		if len(dictionary.Words) != 88 {
			t.Errorf("expected 88 words to be returned but got %d", len(dictionary.Words))
		}

		for _, word := range dictionary.Words {
			if _, ok := dictionary.Map[word]; !ok {
				t.Errorf("expected the word %s to be in the map but it is not", word)
			}
		}
	})
}
