package event

import "time"

type Event interface {
	Name() string
	Id() string
	AggregateId() string
	Data() map[string]interface{}
	CreatedAt() time.Time
}
