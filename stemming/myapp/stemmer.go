package main

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/kljensen/snowball"
	"strings"
)

var apostrophe = "'"

func stemming(words []string) *hashset.Set {
	var stemmedWords = hashset.New()
	for i := 0; i < len(words); i++ {
		stemmedWord, err := snowball.Stem(words[i], "english", true)
		if err == nil && isInvalidWord(stemmedWord) && isWordWithApostropheOk(stemmedWord) {
			stemmedWords.Add(stemmedWord)
		}
	}

	return stemmedWords
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

func isWordWithApostropheOk(str string) bool {
	if !strings.Contains(str, apostrophe) {
		return true
	}
	words := strings.Split(str, apostrophe)
	if wordsAfterApostrophe.Contains(words[1]) {
		return false
	} else {
		return true
	}
}
