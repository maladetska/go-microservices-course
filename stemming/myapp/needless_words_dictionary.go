package main

import "github.com/emirpasic/gods/sets/hashset"

var (
	articles       = hashset.New("a", "an", "the")
	auxiliaryVerbs = hashset.New(
		"am", "are", "be", "been", "being", "can", "could", "do", "does", "did", "had", "have",
		"has", "might", "must", "is", "shall", "should", "was", "were", "will", "would",
	)
	conjunctions = hashset.New(
		"and", "as", "because", "but", "for", "nor", "or", "so", "yet",
	)
	interjection = hashset.New(
		"ahem", "eureka", "ew", "hmm", "hurray", "oh", "ouch", "psst", "shh", "ugh", "uh", "um",
		"um-hum", "wow", "yay", "yippee", "yum",
	)
	prepositions = hashset.New(
		"above", "across", "against", "along", "among", "around", "at", "before", "behind",
		"below", "beneath", "beside", "between", "by", "down", "forward", "from", "in", "into",
		"near", "of", "off", "on", "to", "toward", "under", "up", "upon", "with", "within",
		"without",
	)
	pronouns = hashset.New(
		"all", "another", "any", "anybody", "anyone", "anything", "both", "each", "either",
		"every", "everybody", "everyone", "everything", "he", "her", "hers", "herself", "him",
		"himself", "his", "i", "it", "its", "itself", "me", "mine", "my", "myself", "neither",
		"other", "others", "our", "ours", "ourselves", "she", "some", "somebody", "someone",
		"something", "that", "their", "theirs", "them", "themselves", "these", "they", "this",
		"those", "us", "we", "what", "whatever", "which", "whichever", "who", "whoever", "whose",
		"you", "your", "yours", "yourself", "yourselves",
	)
	quantifiers = hashset.New(
		"enough", "few", "many", "much",
	)
	wordsAfterApostrophe = hashset.New(
		"s", "m", "r", "re", "v", "ve", "d", "ll", "t",
	)
)
