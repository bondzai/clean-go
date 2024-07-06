package repositories

import (
	"clean-go/internal/models"
	"clean-go/pkg/redis"
	"errors"
	"fmt"
)

type redisEventRepo struct {
	client redis.Client
}

func NewRedisEventRepo(client redis.Client) EventRepo {
	return &redisEventRepo{client: client}
}

func (e *redisEventRepo) Create(event *models.Event) error {
	key := fmt.Sprintf("event:%s", event.ID)
	if e.client.Get(key, new(models.Event)) == nil {
		return errors.New("event already exists")
	}

	return e.client.Set(key, event)
}

func (e *redisEventRepo) GetAll() ([]models.Event, error) {
	var events []models.Event
	keys, err := e.client.Keys("event:*")
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		var event models.Event
		if err := e.client.Get(key, &event); err == nil {
			events = append(events, event)
		}
	}

	return events, nil
}

func (e *redisEventRepo) GetById(id string) (*models.Event, error) {
	var event models.Event
	key := fmt.Sprintf("event:%s", id)

	err := e.client.Get(key, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *redisEventRepo) Update(event *models.Event) error {
	key := fmt.Sprintf("event:%s", event.ID)
	if e.client.Get(key, new(models.Event)) != nil {
		return errors.New("event not found")
	}

	return e.client.Set(key, event)
}

func (e *redisEventRepo) Delete(id string) error {
	key := fmt.Sprintf("event:%s", id)
	if e.client.Get(key, new(models.Event)) != nil {
		return errors.New("event not found")
	}

	return e.client.Remove(key)
}
