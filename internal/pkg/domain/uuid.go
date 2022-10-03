package domain

import "github.com/google/uuid"

type Uuid string

func NewUuid() Uuid {
	return Uuid(uuid.New().String())
}

func (u Uuid) String() string {
	return string(u)
}
