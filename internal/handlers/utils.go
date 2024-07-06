package handlers

import (
	"clean-go/pkg/utils/errs"

	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errs.AppError:
		return c.Status(e.Code).JSON(fiber.Map{"error": e.Message})

	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unexpected error"})
	}
}
