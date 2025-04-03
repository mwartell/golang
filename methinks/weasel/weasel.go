package weasel

import (
	"fmt"
)

type Population []string

type Experiment struct {
	target       string
	population   Population
	generation   int
	mutationRate float64
	charset      Charset
}

func NewExperiment(target string, size int, rate float64, charset Charset) (*Experiment, error) {
	exp := &Experiment{
		target:       target,
		population:   newPool(size, len(target), charset),
		generation:   0,
		mutationRate: rate,
		charset:      charset,
	}
	if !exp.validateCharset() {
		return nil, fmt.Errorf("incomplete charset: %s", charset)
	}
	return exp, nil
}

func (e *Experiment) validateCharset() bool {
	// ensure that every character in the target string is in the charset
	for _, c := range e.target {
		if !e.charset[c] {
			return false
		}
	}
	return true
}
