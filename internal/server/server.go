package server

import (
	"github.com/gofiber/fiber/v2"

	"Wordle/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "Wordle",
			AppName:      "Wordle",
		}),

		db: database.New(),
	}

	return server
}
