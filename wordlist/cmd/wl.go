package main

import (
	"fmt"
	"log"

	"mwartell/wordlist"

	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Panic(err)
		return
	}
	cacheDir := env["CACHE_DIR"]

	sources := []string{wordlist.SourceAlpha, wordlist.SourceMedium}

	for _, source := range sources {
		filename, err := wordlist.CacheWords(source, cacheDir)
		if err != nil {
			log.Panic(err)
			return
		}

		wl := wordlist.NewWordList()
		wl.FromFile(filename)
		fmt.Println("Wordlist from", filename)
		fmt.Println(wl[:5])
	}
}
