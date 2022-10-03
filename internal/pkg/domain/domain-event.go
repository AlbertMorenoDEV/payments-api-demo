package domain

import "time"

type DomainEvents []DomainEvent

type DomainEvent interface {
	Name() string
	Id() Uuid
	AggregateId() Uuid
	Data() map[string]interface{}
	CreatedAt() time.Time
}
