package main

import (
	"clean-go/internal/handlers"
	"clean-go/internal/repositories"
	"clean-go/internal/services"
	"clean-go/pkg/redis"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	redisClient, err := redis.NewClient(redis.Config{
		Address:  "localhost:6379",
		Username: "",
		Password: "",
		DB:       0,
	})
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	eventRepo := repositories.NewRedisEventRepo(redisClient)
	eventService := services.NewEventService(eventRepo)
	eventHandler := handlers.NewEventHandler(eventService)

	app := fiber.New()

	v1 := app.Group("/api/v1")

	events := v1.Group("/events")
	events.Post("", eventHandler.JustMiddleWare, eventHandler.CreateEvent)
	events.Get("", eventHandler.GetEvents)
	events.Get("/:id", eventHandler.GetEventById)
	events.Put("/:id", eventHandler.UpdateEvent)
	events.Delete("/:id", eventHandler.DeleteEvent)

	if err := app.Listen(":9000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
