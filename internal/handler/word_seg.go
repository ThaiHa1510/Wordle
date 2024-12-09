package handler

import (
	"Wordle/internal/database"
	"Wordle/internal/response"
	"Wordle/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func WordSegHandler(db *database.Service) func(*fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		var body response.BodyWordsegPost
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}
		validate := validator.New(validator.WithRequiredStructEnabled())

		if err := validate.Struct(&body); err != nil {
			validationErrors := parseValidationErrors(err)
			return c.Status(fiber.StatusUnprocessableEntity).JSON(response.HTTPValidationError{
				Detail: validationErrors,
			})
		}

		err := utils.AddNewWord(body.Text)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Text processed successfully",
		})
	}
}
