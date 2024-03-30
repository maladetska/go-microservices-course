package main

import (
	"flag"
	"fmt"
)

func main() {
	var text string
	flag.StringVar(&text, "s", "different words", "words for stemming")
	flag.Parse()

	var words = splitWords(text)
	var stemmedWords = stemming(words)

	PrintResult(stemmedWords.Values())
}

func PrintResult(set []interface{}) {
	for i := range set {
		fmt.Print(set[i], " ")
	}
}
