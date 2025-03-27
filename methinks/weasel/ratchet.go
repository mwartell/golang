package weasel

import (
	"math/rand"
)

func (e Experiment) RatchetSolve() int {
	attempt := e.charset.RandomString(len(e.target))
	generation := 0

	for attempt != e.target {
		attempt = e.ratchetCloser(attempt)
		generation++
		if generation%100 == 1 || attempt == e.target {
			// fmt.Printf("Generation %d: %s\n", generation, attempt)
		}
	}

	return generation
}

// get a possible step closer by randomly changing a character
// in the attempt string to a character from the charset unless
// it already matches the target
// this is a "ratchet" step
func (e Experiment) ratchetCloser(attempt string) string {
	if attempt == e.target {
		return attempt
	}

	// find indices of mismatched characters
	var mismatches []int
	for i := range attempt {
		if attempt[i] != e.target[i] {
			mismatches = append(mismatches, i)
		}
	}

	// choose a random mismatched index
	mismatchIndex := mismatches[rand.Intn(len(mismatches))]

	// copy characters from attempt to sb, changing the mismatched character
	sb := []rune(attempt)
	sb[mismatchIndex] = e.charset.RandomRune()

	return string(sb)
}
