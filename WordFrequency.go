package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) {
	words := strings.Fields(text)
	freq := make(map[string]int)
	for i, word := range words {
		words[i] = strings.Trim(word, ".!?;',")
		words[i] = strings.ToLower(words[i])
		freq[words[i]]++
	}

	for word, freq := range freq {
		if freq > 1 {
			fmt.Printf("%s %v\n", word, freq)
		}
	}
}

func countWords() {
	text := "A aa b aaa, aaaa a aa. B bb!"
	wordFrequency(text)
}
