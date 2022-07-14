package main

// This code is to find the word in the given sentence.
// The code will first convert the string statement into array and return the instance of the
// word present in the statement.
// The word comparison will be case insensitive.

import (
	"fmt"
	"os"
	"strings"
)

const (
	corpus = "" + "the lazy cat jumps again and again"
)

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Println("Please enter the words to be searched")
		return
	}
	correct := false
	for _, q := range query {
		givenword := strings.ToLower(q)

		for i, w := range words {
			wCase := strings.ToLower(w)
			if wCase == givenword {
				fmt.Printf("#%-2d %s", i, wCase)
				correct = true
				break
			}
		}
	}
	if !correct {
		fmt.Println("Sorry!! None of the word matched to the sentence")
	}
}
