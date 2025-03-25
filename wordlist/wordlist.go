// Package wordlist provides a simple word list data structure for learning golang
package wordlist

import (
	"bufio"
	"os"
	"strings"
)

// WordList represents a collection of words loaded from a file
type WordList []string

// NewWordList creates a new empty WordList
func NewWordList() WordList {
	return WordList{}
}

// LoadFromFile loads words from a file into the WordList
func (wl *WordList) FromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" && !strings.HasPrefix(word, "#") {
			*wl = append(*wl, word)
		}
	}

	return scanner.Err()
}

// Contains checks if a word exists in the WordList
// TODO: this should use a better search algorithm
func (wl WordList) Contains(word string) bool {
	for _, w := range wl {
		if w == word {
			return true
		}
	}
	return false
}

// Add adds a word to the WordList if it doesn't already exist
func (wl *WordList) Add(word string) {
	if !wl.Contains(word) {
		*wl = append(*wl, word)
	}
}

// Remove removes a word from the WordList
func (wl *WordList) Remove(word string) {
	for i, w := range *wl {
		if w == word {
			// Remove the element at index i
			*wl = append((*wl)[:i], (*wl)[i+1:]...)
			return
		}
	}
}

// Size returns the number of words in the WordList
func (wl WordList) Size() int {
	return len(wl)
}

// ToSlice returns all words as a string slice
func (wl WordList) ToSlice() []string {
	// Since WordList is already a slice, just return a copy
	result := make([]string, len(wl))
	copy(result, wl)
	return result
}
