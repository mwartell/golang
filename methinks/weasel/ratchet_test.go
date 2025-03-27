package weasel

import (
	"testing"
)

func TestRatchetCloser_ChangesOneCharacter(t *testing.T) {
	exp := Experiment{
		target:  "methinks it is like a weasel",
		charset: CharsetDefault,
	}

	attempt := "methinks it is like a weasel"
	modified := exp.ratchetCloser(attempt)

	if modified != attempt {
		t.Errorf("Expected no change for a fully matching string, got: %s", modified)
	}

	attempt = "methinks it is like a weasxl"
	modified = exp.ratchetCloser(attempt)

	if len(modified) != len(attempt) {
		t.Fatalf("Modified string length mismatch: expected %d, got %d", len(attempt), len(modified))
	}

	diffCount := 0
	for i := range attempt {
		if attempt[i] != modified[i] {
			diffCount++
		}
	}

	if diffCount != 1 {
		t.Errorf("Expected exactly one character to change, but %d characters changed", diffCount)
	}
}

func TestRatchetCloser_DoesNotModifyMatchingCharacters(t *testing.T) {
	exp := Experiment{
		target:  "methinks it is like a weasel",
		charset: CharsetDefault,
	}

	attempt := "methinks it is like a weasxl"
	modified := exp.ratchetCloser(attempt)

	for i := range attempt {
		if attempt[i] == exp.target[i] && modified[i] != exp.target[i] {
			t.Errorf("Matching character at index %d was incorrectly modified", i)
		}
	}
}

func TestRatchetCloser_EmptyString(t *testing.T) {
	exp := Experiment{
		target:  "",
		charset: CharsetDefault,
	}

	attempt := ""
	modified := exp.ratchetCloser(attempt)

	if modified != "" {
		t.Errorf("Expected empty string, got: %s", modified)
	}
}

func TestRatchetCloser_AllMismatchedCharacters(t *testing.T) {
	exp := Experiment{
		target:  "methinks it is like a weasel",
		charset: CharsetDefault,
	}

	attempt := "xxxxxxxxxxxxxxxxxxxxxxxxxxx"
	modified := exp.ratchetCloser(attempt)

	if len(modified) != len(attempt) {
		t.Fatalf("Modified string length mismatch: expected %d, got %d", len(attempt), len(modified))
	}

	diffCount := 0
	for i := range attempt {
		if attempt[i] != modified[i] {
			diffCount++
		}
	}

	if diffCount != 1 {
		t.Errorf("Expected exactly one character to change, but %d characters changed", diffCount)
	}
}

func TestRatchetCompletion(t *testing.T) {
	target := "all good things"
	charset := FromString(target)

	exp, err := NewExperiment("all good things", 1, 0, charset)
	if err != nil {
		t.Errorf("Failed to create experiment: %v", err)
	}

	generations := exp.RatchetSolve()
	t.Logf("Generations to completion: %d", generations)
}
