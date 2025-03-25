// Package wordlist provides a simple word list data structure for learning golang
package wordlist

import (
	"bufio"
	"os"
	"strings"
)

// WordList represents a collection of words loaded from a file
type WordList []string



// FromFile reads a file containing a list of words and returns a WordList.
// Each line in the file is treated as a word, with leading and trailing
// whitespace trimmed. Lines that are empty or start with a '#' character
// (comments) are ignored.
//
// Parameters:
//   - filename: The path to the file containing the word list.
//
// Returns:
//   - WordList: A slice of strings containing the words from the file.
//   - error: An error if the file cannot be opened or read.
//
// Example usage:
//   wl, err := FromFile("words.txt")
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Println(wl)
func FromFile(filename string) (WordList, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var wl WordList
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
