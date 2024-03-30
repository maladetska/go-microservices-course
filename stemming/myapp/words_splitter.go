package main

import (
	"github.com/emirpasic/gods/sets/hashset"
	"strings"
	"unicode"
)

var (
	punctuationMarks = hashset.New('.', ',', '!', '?', ':', ';', '"', '(', ')', '-')
)

var isCharSeparator = func(char rune) bool {
	return unicode.IsSpace(char) || punctuationMarks.Contains(char)
}

func splitWords(text string) []string {
	return strings.FieldsFunc(text, isCharSeparator)
}
