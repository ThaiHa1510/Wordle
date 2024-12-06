package handler

import (
	"Wordle/internal/response"
	"Wordle/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// func DailyHandler(c *fiber.Ctx) error {
// 	guess := c.Query("guess")
// 	if guess == "" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Query parameter 'guess' is required",
// 		})
// 	}

// 	// Get 'size' with default value 5
// 	sizeStr := c.Query("size", "5")
// 	size, err := strconv.Atoi(sizeStr)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Query parameter 'size' must be an integer",
// 		})
// 	}

// 	// Get 'seed' if provided
// 	seedStr := c.Query("seed", "1")
// 	_, err = strconv.Atoi(seedStr)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Query parameter 'seed' must be an integer",
// 		})
// 	}

// 	if len(guess) != size {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Guess must be the same length as the word",
// 		})
// 	}

// 	// TODO: Implement actual random guessing logic using 'seed'

// 	// For demonstration, we'll return a mock guess result
// 	guessResult := response.GuessResult{
// 		Slot:   1,
// 		Guess:  guess,
// 		Result: "present",
// 	}

// 	return c.Status(fiber.StatusOK).JSON([]response.GuessResult{guessResult})
// }

func DailyHandler(c *fiber.Ctx) error {
	var query GuessQuery

	// Parse query parameters
	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	// Set default size if not provided
	if query.Size == 0 {
		query.Size = 5
	}

	// Validate the query parameters
	if err := guessValidate.Struct(&query); err != nil {
		validationErrors := parseValidationErrors(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.HTTPValidationError{
			Detail: validationErrors,
		})
	}

	// Ensure the guess length matches the size
	if len(query.Guess) != query.Size {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "The length of guess does not match the specified size",
		})
	}

	guessingWord := strings.ToLower(query.Guess)

	// Validate that the guess is a valid word
	if !utils.IsValidWord(guessingWord) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "The guess is not a valid word",
		})
	}

	// Select a random word based on size and seed
	targetWord, err := utils.GetRandomWord(query.Size, query.Seed)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to select a random word",
		})
	}

	// Compare the guess to the target word
	feedback := compareWords(guessingWord, targetWord)

	return c.Status(fiber.StatusOK).JSON(feedback)

}
