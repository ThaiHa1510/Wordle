package server

import (
	"Wordle/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
	s.App.Post("/wordseg", handler.WordSegHandler(&s.db))
	s.App.Get("/daily/", handler.DailyHandler)
	s.App.Get("/word/:word", handler.WordHandler)
	s.App.Get("/random", handler.RandomHandler)

}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}
	return c.JSON(resp)
}
