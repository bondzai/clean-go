package repositories

import (
	"clean-go/internal/models"
	"errors"
)

type eventRepo struct {
	eventMap map[string]models.Event
}

func NewEventRepo() EventRepo {
	return &eventRepo{eventMap: make(map[string]models.Event)}
}

func (e *eventRepo) Create(event *models.Event) error {
	if _, exists := e.eventMap[event.ID]; exists {
		return errors.New("event already exists")
	}
	e.eventMap[event.ID] = *event

	return nil
}

func (e *eventRepo) GetAll() ([]models.Event, error) {
	var events []models.Event
	for _, event := range e.eventMap {
		events = append(events, event)
	}

	return events, nil
}

func (e *eventRepo) GetById(id string) (*models.Event, error) {
	if event, exists := e.eventMap[id]; exists {
		return &event, nil
	}

	return nil, errors.New("event not found")
}

func (e *eventRepo) Update(event *models.Event) error {
	if _, exists := e.eventMap[event.ID]; !exists {
		return errors.New("event not found")
	}
	e.eventMap[event.ID] = *event

	return nil
}

func (e *eventRepo) Delete(id string) error {
	if _, exists := e.eventMap[id]; !exists {
		return errors.New("event not found")
	}
	delete(e.eventMap, id)

	return nil
}
