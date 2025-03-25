package wordlist

import (
	"testing"

	"gotest.tools/v3/fs"
)

func TestWordList(t *testing.T) {
	// Create test file
	file := fs.NewFile(t, "test.txt", fs.WithContent("apple\nbanana\n# comment\ncherry\n\n"))
	defer file.Remove()

	wl := NewWordList()
	err := wl.FromFile(file.Path())
	if err != nil {
		t.Fatalf("Failed to load wordlist: %v", err)
	}

	if wl.Size() != 3 {
		t.Errorf("Expected 3 words, got %d", wl.Size())
	}

	if !wl.Contains("apple") {
		t.Error("Expected 'apple' to be in wordlist")
	}

	if !wl.Contains("banana") {
		t.Error("Expected 'banana' to be in wordlist")
	}

	if !wl.Contains("cherry") {
		t.Error("Expected 'cherry' to be in wordlist")
	}

	if wl.Contains("comment") {
		t.Error("Did not expect 'comment' to be in wordlist")
	}

	wl.Add("date")
	if !wl.Contains("date") {
		t.Error("Expected 'date' to be in wordlist after adding")
	}

	// Test adding duplicate
	oldSize := wl.Size()
	wl.Add("date")
	if wl.Size() != oldSize {
		t.Error("Size should not change when adding duplicate")
	}

	wl.Remove("apple")
	if wl.Contains("apple") {
		t.Error("Did not expect 'apple' to be in wordlist after removing")
	}

	words := wl.ToSlice()
	if len(words) != 3 {
		t.Errorf("Expected 3 words in slice, got %d", len(words))
	}
}

func BenchmarkFromFile(b *testing.B) {
	tmpdir := b.TempDir()
	filename, err := CacheWords(SourceMedium, tmpdir)
	if err != nil {
		b.Fatalf("Failed to download wordlist: %v", err)
	}


	for b.Loop() {
		wl := NewWordList()
		err := wl.FromFile(filename)
		if err != nil {
			b.Fatalf("Failed to load wordlist: %v", err)
		}
	}

}
