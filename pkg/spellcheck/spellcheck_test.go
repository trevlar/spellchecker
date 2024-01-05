package spellcheck

import (
	"bytes"
	"os"
	"spellchecker/internal/dictionary"
	"testing"
)

func handleStdOutAndCheckSpellingForFile(filename string, dict dictionary.Dictionary) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	CheckSpellingForFile(filename, &dict)

	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	buf.ReadFrom(r)

	return buf.String()
}

func TestCheckSpellingForFile(t *testing.T) {
	t.Run("should print out misspelled words when they exist in file", func(t *testing.T) {
		dict := dictionary.Dictionary{
			Words: []string{"hello", "world", "foo", "bar"},
			Map: map[string]bool{
				"hello": true,
				"world": true,
				"foo":   true,
				"bar":   true,
			},
		}
		output := handleStdOutAndCheckSpellingForFile("../../testdata/incorrect-spelling.txt", dict)

		expectedOutput := "========================================\nSummary of Misspelled Words:\n========================================\n\n\x1b[31mbr\x1b[0m, \x1b[31mfo\x1b[0m, \x1b[31mhelo\x1b[0m, \x1b[31mword\x1b[0m\n\n\n========================================\nSpellcheck Details:\n========================================\n\n\n   1. \t\x1b[33mMisspelled Word:\x1b[0m '\x1b[31mbr\x1b[0m'\n\t\x1b[33mOccurrences:\x1b[0m\n\t\ta, \x1b[33mLine\x1b[0m 2, \x1b[33mWord\x1b[0m 2\n\t\t   \x1b[33mContext:\x1b[0m \"... fo \x1b[33mbr\x1b[0m \" \n\t\x1b[33mSuggestions:\x1b[0m \x1b[32mbar\x1b[0m, \x1b[32mfoo\x1b[0m\n\n   2. \t\x1b[33mMisspelled Word:\x1b[0m '\x1b[31mfo\x1b[0m'\n\t\x1b[33mOccurrences:\x1b[0m\n\t\ta, \x1b[33mLine\x1b[0m 2, \x1b[33mWord\x1b[0m 1\n\t\t   \x1b[33mContext:\x1b[0m \"...  \x1b[33mfo\x1b[0m br\" \n\t\x1b[33mSuggestions:\x1b[0m \x1b[32mfoo\x1b[0m, \x1b[32mbar\x1b[0m\n\n   3. \t\x1b[33mMisspelled Word:\x1b[0m '\x1b[31mhelo\x1b[0m'\n\t\x1b[33mOccurrences:\x1b[0m\n\t\ta, \x1b[33mLine\x1b[0m 1, \x1b[33mWord\x1b[0m 1\n\t\t   \x1b[33mContext:\x1b[0m \"...  \x1b[33mhelo\x1b[0m word\" \n\t\x1b[33mSuggestions:\x1b[0m \x1b[32mhello\x1b[0m, \x1b[32mfoo\x1b[0m\n\n   4. \t\x1b[33mMisspelled Word:\x1b[0m '\x1b[31mword\x1b[0m'\n\t\x1b[33mOccurrences:\x1b[0m\n\t\ta, \x1b[33mLine\x1b[0m 1, \x1b[33mWord\x1b[0m 2\n\t\t   \x1b[33mContext:\x1b[0m \"... helo \x1b[33mword\x1b[0m \" \n\t\x1b[33mSuggestions:\x1b[0m \x1b[32mworld\x1b[0m, \x1b[32mfoo\x1b[0m, \x1b[32mbar\x1b[0m\n\n"

		if output != expectedOutput {
			t.Errorf("Expected output to equal %q, got %q", expectedOutput, output)
		}
	})

	t.Run("should not print out misspelled words when all spelled correctly", func(t *testing.T) {
		dict := dictionary.Dictionary{
			Words: []string{"hello", "world", "foo", "bar"},
			Map: map[string]bool{
				"hello": true,
				"world": true,
				"foo":   true,
				"bar":   true,
			},
		}
		output := handleStdOutAndCheckSpellingForFile("../../testdata/correct-spelling.txt", dict)

		if output != "No misspellings found\n" {
			t.Errorf("Expected output to be empty, got %q", output)
		}
	})
}
