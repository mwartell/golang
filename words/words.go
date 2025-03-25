// Package words provides a simple word data structure for learning golang
package words

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

// Words represents a collection of words
type Words []string



// FromFile reads a file containing a words and returns a Words.
// Each line in the file is treated as a word, with leading and trailing
// whitespace trimmed. Lines that are empty or start with a '#' character
// (comments) are ignored.
//
// Parameters:
//   - filename: The path to the file containing the words.
//
// Returns:
//   - Words: a collection of words read from the file.
//   - error: An error if the file cannot be opened or read.
//
func FromFile(filename string) (Words, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var wl Words
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" && !strings.HasPrefix(word, "#") {
			wl = append(wl, word)
		}
	}

	return wl, scanner.Err()
}

// Contains checks if a word exists in the WordList
func (wl Words) Contains(word string) bool {
	return slices.Contains(wl, word)
}

// Add adds a word to the WordList if it doesn't already exist
func (wl *Words) Add(word string) {
	if !wl.Contains(word) {
		*wl = append(*wl, word)
	}
}

// Remove removes a word from the WordList
func (wl *Words) Remove(word string) {
	for i, w := range *wl {
		if w == word {
			// Remove the element at index i
			*wl = append((*wl)[:i], (*wl)[i+1:]...)
			return
		}
	}
}

// Size returns the number of words in the WordList
func (wl Words) Size() int {
	return len(wl)
}
