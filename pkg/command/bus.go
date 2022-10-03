package command

import (
	"context"
	"fmt"
	"sync"
)

type Bus interface {
	Register(command Command, handler Handler) error
	Dispatch(ctx context.Context, command Command) error
}

type CommandBus struct {
	handlers map[string]Handler
	lock     sync.Mutex
	wg       sync.WaitGroup
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]Handler),
		lock:     sync.Mutex{},
		wg:       sync.WaitGroup{},
	}
}

func (bus *CommandBus) Register(command Command, handler Handler) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	if _, ok := bus.handlers[command.CommandName()]; ok {
		return fmt.Errorf("command <%s> already registered", command.CommandName())
	}

	bus.handlers[command.CommandName()] = handler

	return nil
}

func (bus *CommandBus) Dispatch(ctx context.Context, command Command) error {
	if handler, ok := bus.handlers[command.CommandName()]; ok {
		return handler.Handle(ctx, command)
	}

	return fmt.Errorf("command <%s> not registered", command.CommandName())
}
