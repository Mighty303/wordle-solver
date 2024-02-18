package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// LoadWords would load the list of valid words from the database or a file.
func LoadWords() []string {
	dat, err := os.ReadFile("words.txt")
	check(err)
	// fmt.Println(string(dat))

	// Placeholder: Load words from a database or file.
	// words := strings.Split(string(dat), " ")
	// fmt.Println(words)
	return strings.Split(string(dat), "\n")
}

// FilterWords filters the list of words based on the game's feedback.
func FilterWords(words []string, feedback string) []string {
	// Placeholder: Implement the filtering logic based on the feedback.
	// For example, if the feedback indicates that the letter 'a' is in the correct position 1,
	// you would filter the list to include only words with 'a' in position 1.

	return words // Return the filtered list of words.
}

// SuggestWord suggests a word from the list of possible words.
func SuggestWord(words []string) string {
	// Placeholder: Implement the logic to choose the next word to suggest.
	// This could be based on frequency analysis, elimination strategy, or random selection.
	return words[0] // Return the first word in the list as a simple example.
}

func main() {
	// Load the list of valid words.
	words := LoadWords()

	// Placeholder: Loop to receive user input and provide suggestions.
	feedback := "some feedback from the game"
	words = FilterWords(words, feedback)
	suggestion := SuggestWord(words)
	fmt.Printf("Suggested word: %s\n", suggestion)
}
