package weasel

import (
	"math/rand"
)

func PoolSolve(e *Experiment) int {
	for {
		e.generation++
		bestMatch := e.population.getBestMatch(e.target)

		// log.Printf("Generation %d: %s\n", e.generation, bestMatch)
		if bestMatch == e.target {
			return e.generation
		}

		e.nextGeneration(bestMatch)
	}
}

func newPool(size, length int, charset Charset) Population {
	population := make(Population, size)
	for i := range population {
		population[i] = charset.RandomString(length)
	}
	return population
}

func (population Population) getBestMatch(target string) string {
	bestMatch := population[0]
	bestScore := score(bestMatch, target)

	for _, individual := range population {
		currentScore := score(individual, target)
		if currentScore > bestScore {
			bestMatch = individual
			bestScore = currentScore
		}
	}
	return bestMatch
}

func score(candidate, target string) int {
	score := 0
	for i := range candidate {
		if candidate[i] == target[i] {
			score++
		}
	}
	return score
}

// the next generation is created replacing the current population with new individuals
// that are mutated versions of the best match
func (e Experiment) nextGeneration(parent string) {
	for i := range e.population {
		e.population[i] = e.mutantChild(parent)
	}
}

// mutantChild creates a new individual by mutating the parent string
// by replacing random characters with random characters from the charset
// the mutation rate determines the probability of mutation for each character
func (e Experiment) mutantChild(parent string) string {
	mut := []rune(parent)
	for i := range mut {
		if rand.Float64() < e.mutationRate {
			mut[i] = e.charset.RandomRune()
		}
	}
	return string(mut)
}
