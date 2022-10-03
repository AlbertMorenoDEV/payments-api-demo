package infrastructure

import "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"

type ChannelDomainEventPublisher struct {
	domainEvents chan domain.DomainEvent
}

func NewChannelDomainEventPublisher(channel chan domain.DomainEvent) *ChannelDomainEventPublisher {
	return &ChannelDomainEventPublisher{domainEvents: channel}
}

func (p ChannelDomainEventPublisher) Publish(events domain.DomainEvents) {
	for _, event := range events {
		p.domainEvents <- event
	}
}
