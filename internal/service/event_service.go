package service

import (
	"errors"
	"github.com/EnesHarman/eventify/internal/model"
	"github.com/EnesHarman/eventify/internal/repository"
	"github.com/labstack/gommon/log"
)

type EventService interface {
	InsertEvent(event model.Event) error
	GetEvents(page, size int64) ([]model.Event, error)
}

type EventServiceImpl struct {
	eventRepository repository.EventRepository
}

func (service EventServiceImpl) InsertEvent(event model.Event) error {
	err := service.eventRepository.InsertEvent(event)
	if err != nil {
		return errors.New("Error while inserting event")
	}
	return err
}

func (service EventServiceImpl) GetEvents(page, size int64) ([]model.Event, error) {
	if page < 1 || size < 1 {
		log.Error("Page and size should be greater than 1")
		return nil, errors.New("Page and size should be greater than 1")

	}
	events, err := service.eventRepository.GetEvents(page, size)
	if err != nil {
		log.Error("Error while getting events %v", err)
		return nil, errors.New("Error while getting events")
	}
	return events, err
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &EventServiceImpl{
		eventRepository: eventRepository,
	}
}
