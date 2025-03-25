package main

import (
	"fmt"
	"log"

	"mwartell/words"

	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Panic(err)
		return
	}
	cacheDir := env["CACHE_DIR"]

	sources := []string{words.SourceAlpha, words.SourceMedium}

	for _, source := range sources {
		filename, err := words.CacheWords(source, cacheDir)
		if err != nil {
			log.Panic(err)
			return
		}

		wl, err := words.FromFile(filename)
		if err != nil {
			log.Println()
		} else {
			fmt.Printf("Wordlist from %v contains %d words\n", filename, wl.Size())
			fmt.Println(wl[:5])
		}
	}
}
