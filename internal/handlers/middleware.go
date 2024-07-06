package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h *eventHandler) JustMiddleWare(c *fiber.Ctx) error {
	log.Println("just a simple middleware...")
	return c.Next()
}
