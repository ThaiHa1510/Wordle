// utils/wordlist.go
package utils

import (
	"bufio"
	"embed"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
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

func GetRandomWord(size int, seed int64) (string, error) {
	var filteredWords []string
	for _, word := range wordList {
		if len(word) == size {
			filteredWords = append(filteredWords, word)
		}
	}

	if len(filteredWords) == 0 {
		return "", errors.New("no words found with the specified size")
	}

	if seed >= 0 && int(seed) < len(filteredWords) {
		return filteredWords[seed], nil
	}
	var rng *rand.Rand
	if seed != 0 {
		rng = rand.New(rand.NewSource(seed))
	} else {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	randomIndex := rng.Intn(len(filteredWords))
	if randomIndex < 0 || randomIndex >= len(filteredWords) {
		randomIndex = 0
	} else {
		return filteredWords[randomIndex], nil
	}

	return filteredWords[randomIndex], nil
}

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
	file, err := dailyFile.Open("daily.txt")
	if err != nil {
		panic("Failed to open daily.txt: " + err.Error())
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

func GetDailyWord(size int, seed int64) (string, error) {
	var rng *rand.Rand
	if seed != 0 {
		rng = rand.New(rand.NewSource(seed))
	} else {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	var filteredWords []string
	for _, word := range dailyList {
		if len(word) == size {
			filteredWords = append(filteredWords, word)
		}
	}

	if len(filteredWords) == 0 {
		return "", errors.New("no words found with the specified size")
	}

	randomIndex := rng.Intn(len(filteredWords))
	return filteredWords[randomIndex], nil
}

// AddNewWord adds a new word to words.txt and updates the in-memory wordList.
// It ensures that the word is not already present and is alphabetic.
func AddNewWord(newWord string) error {
	newWord = strings.TrimSpace(strings.ToLower(newWord))
	if newWord == "" {
		return errors.New("word cannot be empty")
	}

	if !isAlphabetic(newWord) {
		return errors.New("word must contain only alphabetic characters")
	}

	for _, word := range wordList {
		if word == newWord {
			return errors.New("word already exists in the list")
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		return errors.New("failed to open words.txt: " + err.Error())
	}
	filePath := filepath.Join(pwd, "../../internal/utils/words.txt")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("failed to open words.txt: " + err.Error())
	}
	defer file.Close()

	if _, err := file.WriteString(newWord + "\n"); err != nil {
		return errors.New("failed to write to words.txt: " + err.Error())
	}

	wordList = append(wordList, newWord)
	return nil
}

// isAlphabetic checks if a string contains only alphabetic characters.
func isAlphabetic(s string) bool {
	for _, r := range s {
		if !('a' <= r && r <= 'z') && !('A' <= r && r <= 'Z') {
			return false
		}
	}
	return true
}
