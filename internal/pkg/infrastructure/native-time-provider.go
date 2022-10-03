package infrastructure

import "time"

type SystemTimeProvider struct {
}

func (dp SystemTimeProvider) Now() time.Time {
	date := time.Now()

	return date
}
