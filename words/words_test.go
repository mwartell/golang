package words

import (
	"fmt"
	"log"
	"testing"

	"gotest.tools/v3/fs"
)

func TestWordList(t *testing.T) {
	// Create test file
	file := fs.NewFile(t, "test.txt", fs.WithContent("apple\nbanana\n# comment\ncherry\n\n"))
	defer file.Remove()

	wl, err := FromFile(file.Path())
	if err != nil {
		t.Fatalf("Failed to load wordlist: %v", err)
	}

	if wl.Size() != 3 {
		t.Errorf("Expected 3 words, got %d", wl.Size())
	}

	presences := map[string]bool{
		"apple":     true,
		"banana":    true,
		"cherry":    true,
		"date":      false,
		"comment":   false,
		"# comment": false,
	}
	for target, expected := range presences {
		presence := map[bool]string{true: "present", false: "absent"}[expected]
		if wl.Contains(target) != expected {
			t.Fatalf("Expected %v to be %v", target, presence)
		}
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
}

func BenchmarkFromFile(b *testing.B) {
	tmpdir := b.TempDir()
	filename, err := CacheWords(SourceMedium, tmpdir)
	if err != nil {
		b.Fatalf("Failed to download wordlist: %v", err)
	}

	for b.Loop() {
		_, err := FromFile(filename)
		if err != nil {
			b.Fatalf("Failed to load wordlist: %v", err)
		}
	}
}

func ExampleWords() {
	wl, err := FromFile("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wl.Size())
	// Output:
	// 104334
}
