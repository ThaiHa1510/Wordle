package handler

import (
	"Wordle/internal/utils"

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

	results := utils.CompareWords(guess, word)

	return c.Status(fiber.StatusOK).JSON(results)
}
