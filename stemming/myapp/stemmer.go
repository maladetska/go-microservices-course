package main

import (
	"github.com/kljensen/snowball"
	"strings"
)

type HashSet map[string]bool

var apostrophe = "'"

func stemming(words []string) []string {
	var stemmedWords = make(HashSet)
	for _, word := range words {
		stemmedWord, err := snowball.Stem(word, "english", true)
		if err == nil &&
			isInvalidWord(stemmedWord) &&
			isWordWithApostropheAppropriate(stemmedWord) {
			stemmedWords[stemmedWord] = true
		}
	}

	return getWordSlice(&stemmedWords)
}

func isInvalidWord(str string) bool {
	if articles.Contains(str) || auxiliaryVerbs.Contains(str) ||
		conjunctions.Contains(str) || interjection.Contains(str) ||
		prepositions.Contains(str) || pronouns.Contains(str) ||
		quantifiers.Contains(str) {
		return false
	}
	return true
}

func isWordWithApostropheAppropriate(str string) bool {
	if !strings.Contains(str, apostrophe) {
		return true
	}
	words := strings.Split(str, apostrophe)

	return !wordsAfterApostrophe.Contains(words[1])
}

func getWordSlice(wordsSet *HashSet) []string {
	words := make([]string, 0, len(*wordsSet))
	for k := range *wordsSet {
		words = append(words, k)
	}

	return words
}
