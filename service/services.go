package service

import (
	"context"
	"errors"
	"log"

	"github.com/tigerinus/xin/repository"
)

type Services struct {
	EventService  *EventService
	ActionService *ActionService
}

var (
	logger = log.Default()

	ErrInboundChannelNotFound     = errors.New("inbound channel not found")
	ErrSubscriberChannelsNotFound = errors.New("subscriber channels not found")
	ErrAlreadySubscribed          = errors.New("already subscribed")
)

func (s *Services) Start(ctx *context.Context) {
	go s.EventService.Start(ctx)
	go s.ActionService.Start(ctx)
}

func NewServices(repository *repository.Repository) Services {
	return Services{
		EventService:  NewEventService(repository),
		ActionService: NewActionService(repository),
	}
}
