package handlers

import (
	"clean-go/internal/models"
	"clean-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

type EventHandler interface {
	CreateEvent(c *fiber.Ctx) error
	GetEvents(c *fiber.Ctx) error
	GetEventById(c *fiber.Ctx) error
	UpdateEvent(c *fiber.Ctx) error
	DeleteEvent(c *fiber.Ctx) error

	//middleware
	JustMiddleWare(c *fiber.Ctx) error
}

type eventHandler struct {
	service services.EventService
}

func NewEventHandler(service services.EventService) EventHandler {
	return &eventHandler{service: service}
}

func (h *eventHandler) CreateEvent(c *fiber.Ctx) error {
	var event models.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := h.service.Create(&event); err != nil {
		return handleError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "Event created successfully",
	})
}

func (h *eventHandler) GetEvents(c *fiber.Ctx) error {
	data, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get events",
		})
	}

	return c.JSON(data)
}

func (h *eventHandler) GetEventById(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := h.service.GetById(id)
	if err != nil {
		return handleError(c, err)
	}

	return c.JSON(data)
}

func (h *eventHandler) UpdateEvent(c *fiber.Ctx) error {
	var event models.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	id := c.Params("id")
	event.ID = id

	if err := h.service.Update(&event); err != nil {
		return handleError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "Event updated successfully",
	})
}

func (h *eventHandler) DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(id); err != nil {
		return handleError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "Event deleted successfully",
	})
}
