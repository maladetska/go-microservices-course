package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	textPtr := flag.String("s", "different words", "words for stemming")
	flag.Parse()

	var words = strings.Fields(*textPtr)
	var stemmedWords = stemming(words)

	PrintResult(stemmedWords.Values())
}

func PrintResult(set []interface{}) {
	for i := range set {
		fmt.Print(set[i], " ")
	}
}
