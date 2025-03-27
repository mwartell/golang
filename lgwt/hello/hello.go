package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const english = "English"

var greetingLookup = map[string]string{
	spanish: "Hola, ",
	french:  "Bonjour, ",
	english: "Hello, ",
	"":      "Hello, ",
}

// Hello returns a personalized greeting in a given language.
func Hello(name string, language string) string {
	return greetingLookup[language] + name
}

func main() {
	fmt.Println(Hello("world", english))
}
