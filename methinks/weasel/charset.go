package weasel

import (
	"math/rand"
	"slices"
)

type Charset map[rune]bool

var CharsetDefault = FromString("abcdefghijklmnopqrstuvwxyz ")

func FromString(s string) Charset {
	seen := make(Charset)
	for _, c := range s {
		seen[c] = true
	}
	return seen
}

// String returns a string representation of the charset
// by sorting the keys in ascending order
func (c Charset) String() string {
	keys := generateRuneSlice(c)
	slices.Sort(keys)
	return string(keys)
}

// RandomString generates a random string of the given length
// using the characters in the charset
func (c Charset) RandomString(length int) string {
	keys := generateRuneSlice(c)
	result := make([]rune, length)

	for i := 0; i < length; i++ {
		result[i] = keys[rand.Intn(len(keys))]
	}
	return string(result)
}

// RandomRune returns a random rune from the Charset
func (c Charset) RandomRune() rune {
	keys := generateRuneSlice(c)
	return keys[rand.Intn(len(keys))]
}

func generateRuneSlice(c Charset) []rune {
	keys := make([]rune, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}
