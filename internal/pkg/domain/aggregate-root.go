package domain

type AggregateRoot interface {
	PullDomainEvents() DomainEvents
}
