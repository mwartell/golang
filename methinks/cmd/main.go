package main

import (
	"fmt"
	"methinks/weasel"
)

func main() {
	cs := weasel.FromString("hello")
	fmt.Printf("cs: %v\n", cs)
}
