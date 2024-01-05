package main

import (
	"os/exec"
	"testing"
)

func handleSpellcheckerRequest(dictionaryFileName string, fileToCheckFileName string) (string, error) {

	cmd := exec.Command("go", "run", "../../cmd/spellchecker", dictionaryFileName, fileToCheckFileName)

	result, err := cmd.CombinedOutput()

	return string(result), err
}

func TestSpellcheckerCLI(t *testing.T) {
	testDictionary := "../../testdata/dictionary.txt"
	testFile := "../../testdata/correct-spelling.txt"

	output, err := handleSpellcheckerRequest(testDictionary, testFile)
	if err != nil {
		t.Fatalf("Error executing command: %v", err)
	}

	expectedOutput := "No misspellings found\n"
	if string(output) != expectedOutput {
		t.Fatalf("Unexpected output: %q, got '%q'", output, expectedOutput)
	}
}
