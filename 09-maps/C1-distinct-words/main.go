package main

import "strings"

func countDistinctWords(messages []string) int {
	distinctWords := make(map[string]bool)

	for _, message := range messages {
		words := strings.Fields(message)
		for _, word := range words {
			lowercaseWord := strings.ToLower(word)
			distinctWords[lowercaseWord] = true
		}
	}
	return len(distinctWords)
	}
