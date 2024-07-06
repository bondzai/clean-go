package services

import (
	"clean-go/internal/models"
	"clean-go/internal/repositories"
)

type EventService interface {
	Create(event *models.Event) error
	GetAll() ([]models.Event, error)
	GetById(id string) (*models.Event, error)
	Update(event *models.Event) error
	Delete(id string) error
}

type eventService struct {
	eventRepo repositories.EventRepo
}

func NewEventService(eventRepo repositories.EventRepo) EventService {
	return &eventService{
		eventRepo: eventRepo,
	}
}

func (s *eventService) Create(event *models.Event) error {
	err := s.eventRepo.Create(event)
	if err != nil {
		return handleDbError(err)
	}

	return nil
}

func (s *eventService) GetAll() ([]models.Event, error) {
	events, err := s.eventRepo.GetAll()
	if err != nil {
		return []models.Event{}, handleDbError(err)
	}

	if events == nil {
		return []models.Event{}, nil
	}

	return events, nil
}

func (s *eventService) GetById(id string) (*models.Event, error) {
	event, err := s.eventRepo.GetById(id)
	if err != nil {
		return nil, handleDbError(err)
	}

	return event, nil
}

func (s *eventService) Update(event *models.Event) error {
	err := s.eventRepo.Update(event)
	if err != nil {
		return handleDbError(err)
	}

	return nil
}

func (s *eventService) Delete(id string) error {
	err := s.eventRepo.Delete(id)
	if err != nil {
		return handleDbError(err)
	}

	return nil
}
