package dictionary

import (
	"bufio"
	"os"
)

type Dictionary struct {
	Words []string
	Map   map[string]bool
}

func ReadFile(filename string) (*Dictionary, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dictionary := &Dictionary{
		Words: make([]string, 0),
		Map:   make(map[string]bool),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		dictionary.Words = append(dictionary.Words, word)
		dictionary.Map[word] = true
	}

	return dictionary, scanner.Err()
}
