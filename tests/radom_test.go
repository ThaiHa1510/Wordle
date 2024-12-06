// handlers/random_test.go
package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRandomHandler(t *testing.T) {
	// Initialize Fiber
	app := fiber.New()

	// Register the route
	app.Get("/api/guess", RandomHandler)

	// Prepare test cases
	testCases := []struct {
		name           string
		query          string
		userID         string
		expectedStatus int
	}{
		{
			name:           "Successful Guess",
			query:          "guess=apple&size=5&seed=1",
			userID:         "1",
			expectedStatus: fiber.StatusOK,
		},
		{
			name:           "Missing Guess",
			query:          "size=5&seed=1",
			userID:         "1",
			expectedStatus: fiber.StatusUnprocessableEntity,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/guess?"+tc.query, nil)
			req.Header.Set("X-User-ID", tc.userID)
			resp, _ := app.Test(req, -1)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode)
		})
	}
}
