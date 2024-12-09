// server/server_test.go
package server

import (
	"encoding/json"
	"testing"

	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type MockDB struct{}

// TestHelloWorldHandler tests the '/' endpoint.
func TestHelloWorldHandler(t *testing.T) {
	// Initialize Fiber app
	app := fiber.New()

	server := &FiberServer{
		App: app,
		db:  nil,
	}

	server.RegisterFiberRoutes()

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var responseBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.NoError(t, err)

	expected := map[string]string{
		"message": "Hello World",
	}

	assert.Equal(t, expected, responseBody)
}
