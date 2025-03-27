// Title: Weasel Program in Go
// Description: A simple genetic algorithm that evolves a population of strings towards a target string.

package weasel

import (
	"math/rand"
	"strings"
)

func Weasel1() int {
	attempt := RandomString(len(Target))
	generation := 0

	for attempt != Target {
		attempt = closer(attempt, Target)
		generation++
		if generation%100 == 1 || attempt == Target {
			// fmt.Printf("Generation %d: %s\n", generation, attempt)
		}
	}

	return generation
}

// mutate returns a new attempt by randomly changing one random character
// that doesn't already match the target
func closer(attempt, target string) string {
	// find indices of mismatched characters
	var mismatches []int
	for i := range attempt {
		if attempt[i] != target[i] {
			mismatches = append(mismatches, i)
		}
	}

	// choose a random mismatched index
	mismatchIndex := mismatches[rand.Intn(len(mismatches))]

	// copy characters from attempt to sb, changing the mismatched character
	var sb strings.Builder
	for i := range attempt {
		if i == mismatchIndex {
			sb.WriteByte(Charset[rand.Intn(len(Charset))])
		} else {
			sb.WriteByte(attempt[i])
		}
	}

	return sb.String()
}
