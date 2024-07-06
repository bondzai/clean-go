package repositories

import "clean-go/internal/models"

type EventRepo interface {
	Create(event *models.Event) error
	GetAll() ([]models.Event, error)
	GetById(id string) (*models.Event, error)
	Update(event *models.Event) error
	Delete(id string) error
}
