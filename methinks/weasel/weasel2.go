package weasel

import (
	"math/rand"
	"strings"
)

func Weasel2() int {
	populationSize := 100
	mutationRate := 0.05

	population := initializePopulation(populationSize, len(Target))
	generation := 0

	for {
		generation++
		bestMatch := getBestMatch(population)
		// fmt.Printf("Generation %d: %s\n", generation, bestMatch)

		if bestMatch == Target {
			return generation
		}

		population = nextGeneration(bestMatch, population, mutationRate)
	}
}

func initializePopulation(size, length int) []string {
	population := make([]string, size)
	for i := range population {
		population[i] = RandomString(length)
	}
	return population
}

func getBestMatch(population []string) string {
	bestMatch := population[0]
	bestScore := score(bestMatch)

	for _, individual := range population {
		currentScore := score(individual)
		if currentScore > bestScore {
			bestMatch = individual
			bestScore = currentScore
		}
	}
	return bestMatch
}

func score(candidate string) int {
	score := 0
	for i := range candidate {
		if candidate[i] == Target[i] {
			score++
		}
	}
	return score
}

// the next generation is created by mutating the best match
func nextGeneration(bestMatch string, population []string, mutationRate float64) []string {
	next := make([]string, len(population))
	for i := range population {
		next[i] = mutate(bestMatch, mutationRate)
	}
	return next
}

func mutate(individual string, mutationRate float64) string {
	var sb strings.Builder
	for i := range individual {
		if rand.Float64() < mutationRate {
			sb.WriteByte(Charset[rand.Intn(len(Charset))])
		} else {
			sb.WriteByte(individual[i])
		}
	}
	return sb.String()
}
