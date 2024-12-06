package handler

import (
	"Wordle/internal/response"

	"github.com/gofiber/fiber/v2"
)

func WordHandler(c *fiber.Ctx) error {
	word := c.Params("word")
	guess := c.Query("guess")
	if guess == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Query parameter 'guess' is required",
		})
	}
	if len(word) != len(guess) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Guess must be the same length as the word",
		})
	}

	results := make([]response.GuessResult, len(guess))
	wordMap := make(map[rune]int)

	// Populate a map with counts of each character in the word.
	for _, char := range word {
		wordMap[char]++
	}

	for i, char := range guess {
		result := "absent"
		if rune(word[i]) == char {
			result = "correct"
			wordMap[char]--
		}
		results[i] = response.GuessResult{
			Slot:   i,
			Guess:  string(char),
			Result: result,
		}
	}

	for i := range guess {
		if results[i].Result == "correct" {
			continue
		}

		char := rune(guess[i])
		if wordMap[char] > 0 {
			results[i].Result = "present"
			wordMap[char]--
		}
	}

	return c.Status(fiber.StatusOK).JSON(results)
}
