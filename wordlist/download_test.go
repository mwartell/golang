package wordlist

import "testing"

func TestCacheWords(t *testing.T) {
	// Create a new WordList

	tmpdir := t.TempDir()

	// Download the wordlist from the internet
	filename, err := CacheWords(SourceMedium, tmpdir)
	if err != nil {
		t.Fatalf("Failed to download wordlist: %v", err)
	}

	if !fileExists(filename) {
		t.Errorf("File %s was not downloaded.", filename)
	}

	_, err = CacheWords("https://not-really-a-domain-foof.com/bogus.txt", tmpdir)
	if err == nil {
		t.Error("Expected an error, got nil")
	}

	_, err = CacheWords("1234", tmpdir)
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
