// utils/wordlist.go
package utils

import (
	"bufio"
	"embed"
	"errors"
	"math/rand"
	"strings"
	"time"
)

//go:embed words.txt
var wordFile embed.FS

var wordList []string

//go:embed daily.txt
var dailyFile embed.FS

var dailyList []string

func init() {
	LoadWords()
	LoadDailyWords()
}

// LoadWords loads words from the embedded words.txt file.
// Ensure that words.txt is placed in the utils directory.
func LoadWords() {
	file, err := wordFile.Open("words.txt")
	if err != nil {
		panic("Failed to open words.txt: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			wordList = append(wordList, strings.ToLower(word))
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Failed to read words.txt: " + err.Error())
	}
}

// GetRandomWord selects a random word of the given size.
// If seed is provided (non-zero), it initializes the random generator with the seed.
func GetRandomWord(size int, seed int64) (string, error) {
	var rng *rand.Rand
	if seed != 0 {
		rng = rand.New(rand.NewSource(seed))
	} else {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	// Filter words by size
	var filteredWords []string
	for _, word := range wordList {
		if len(word) == size {
			filteredWords = append(filteredWords, word)
		}
	}

	if len(filteredWords) == 0 {
		return "", errors.New("no words found with the specified size")
	}

	// Select a random word from the filtered list
	randomIndex := rng.Intn(len(filteredWords))
	return filteredWords[randomIndex], nil
}

// IsValidWord checks if a word exists in the word list.
func IsValidWord(word string) bool {
	word = strings.ToLower(word)
	for _, w := range wordList {
		if w == word {
			return true
		}
	}
	return false
}

func LoadDailyWords() {
	file, err := wordFile.Open("daily.txt")
	if err != nil {
		panic("Failed to open words.txt: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			dailyList = append(dailyList, strings.ToLower(word))
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Failed to read words.txt: " + err.Error())
	}
}

// GetDailyWord selects a random word of the given size.
// If seed is provided (non-zero), it initializes the random generator with the seed.
func GetDailyWord(size int, seed int64) (string, error) {
	var rng *rand.Rand
	if seed != 0 {
		rng = rand.New(rand.NewSource(seed))
	} else {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	// Filter words by size
	var filteredWords []string
	for _, word := range dailyList {
		if len(word) == size {
			filteredWords = append(filteredWords, word)
		}
	}

	if len(filteredWords) == 0 {
		return "", errors.New("no words found with the specified size")
	}

	// Select a random word from the filtered list
	randomIndex := rng.Intn(len(filteredWords))
	return filteredWords[randomIndex], nil
}
