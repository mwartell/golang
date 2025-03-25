package wordlist

import "testing"

const GITHUB_URL = "https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt"

// TestDownload tests the download functionality of the WordList
// with **side effects** (it downloads a file from the internet
// and writes it to disk)
func TestCacheWords(t *testing.T) {
	// Create a new WordList

	// Download the wordlist from the internet
	filename, _ := cacheWords(GITHUB_URL)

	// Check if the file was downloaded
	if !fileExists(filename) {
		t.Errorf("File %s was not downloaded.", filename)
	}

	_, err := cacheWords("https://not-really-a-domain-foof.com/bogus.txt")
	if err == nil {
		t.Error("Expected an error, got nil")
	}

	_, err = cacheWords("1234")
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
