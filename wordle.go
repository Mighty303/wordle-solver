package main

import (
	"fmt"
	"os"
	"strings"
)

// Check if theres an error reading the file
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
	words := strings.Split(string(dat), "\n")
	return words
}

// FilterWords filters the list of words based on the game's feedback.
func FilterWords(words []string, correct string, valid string, absent string) []string {
	newWords := make([]string, 0, len(words))

	// For example, if the feedback indicates that the letter 'a' is in the correct position 1,
	// you would filter the list to include only words with 'a' in position 1.
	correctPositions := strings.Split(correct, "")
	absentPositions := strings.Split(absent, "")

	// Eliminate non correct words
	for _, word := range words {
		if len(word) > 0 {
			match := true
			for i := 0; i < len(correctPositions); i++ {
				chr := strings.ToLower(correctPositions[i])
				if chr != "." && chr != string(rune(word[i])) {
					match = false
				}
			}

			for i := 0; i < len(absentPositions); i++ {
				if strings.Contains(word, strings.ToLower(absentPositions[i])) {
					match = false
				}
			}

			if match {
				newWords = append(newWords, word)
				fmt.Printf("%s\n", word)
			}
		}
	}

	// fmt.Println(correctPositions)
	return newWords // Return the filtered list of words.
}

// SuggestWord suggests a word from the list of possible words.
func SuggestWord(words []string) string {
	// Placeholder: Implement the logic to choose the next word to suggest.
	// This could be based on frequency analysis, elimination strategy, or random selection.
	if len(words) > 0 {
		return words[0] // Return the first word in the list as a simple example.
	} else {
		return "No suggestions available" // Return a message indicating no suggestions.
	}
}

func main() {
	var correct string
	var valid string
	var absent string

	// Load the list of valid words.
	var words = LoadWords()

	// Placeholder: Loop to receive user input and provide suggestions.
	// fmt.Print("Enter correct letters (ex. A..L.): ")
	// fmt.Scan(&correct)
	// fmt.Print("Enter valid letters (ex. PE): ")
	// fmt.Scan(&valid)
	// fmt.Print("Enter absent letters: (ex. DVBNM): ")
	// fmt.Scan(&absent)

	correct = "A..L."
	valid = "PE"
	absent = "DVBMNZX"

	words = FilterWords(words, correct, valid, absent)
	suggestion := SuggestWord(words)
	fmt.Printf("Suggested words: %s\n", suggestion)
}
