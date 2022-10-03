package command

import "context"

type Handler interface {
	Handle(ctx context.Context, command Command) error
}
