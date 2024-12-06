package handler

import (
	"Wordle/internal/database"
	"Wordle/internal/models"
	"Wordle/internal/response"

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
		user := c.Locals("user").(models.User)

		// Create word records
		var words []models.Word
		word := models.Word{
			Content: body.Text,
			UserID:  user.ID,
		}
		words = append(words, word)

		// Save words to the database
		// if err := db.Create(&words).Error; err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"error": "Failed to store words",
		// 	})
		// }

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Text processed successfully",
		})
	}
}
