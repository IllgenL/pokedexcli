package main

import "strings"

func cleanInput(text string) []string {
	words := strings.Fields(text)
	cleanedWords := make([]string, len(words))
	for i, word := range words {
		cleanedWords[i] = strings.ToLower(word)
	}
	return cleanedWords
}
