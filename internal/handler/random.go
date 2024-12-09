package handler

import (
	"Wordle/internal/response"

	"Wordle/internal/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type GuessQuery struct {
	Guess string `query:"guess" validate:"required"`
	Size  int    `query:"size" validate:"omitempty,min=3,max=15"`
	Seed  int64  `query:"seed" validate:"omitempty"`
}

var guessValidate = validator.New()

func RandomHandler(c *fiber.Ctx) error {
	var query GuessQuery

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	if query.Size == 0 {
		query.Size = 5
	}

	if err := guessValidate.Struct(&query); err != nil {
		validationErrors := parseValidationErrors(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.HTTPValidationError{
			Detail: validationErrors,
		})
	}

	if len(query.Guess) != query.Size {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "The length of guess does not match the specified size",
		})
	}

	guessingWord := strings.ToLower(query.Guess)

	if !utils.IsValidWord(guessingWord) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "The guess is not a valid word",
		})
	}

	targetWord, err := utils.GetRandomWord(query.Size, query.Seed)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to select a random word",
		})
	}
	feedback := utils.CompareWords(guessingWord, targetWord)

	return c.Status(fiber.StatusOK).JSON(feedback)

}
