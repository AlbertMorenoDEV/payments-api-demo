package domain

type DomainEventPublisher interface {
	Publish(events DomainEvents)
}
