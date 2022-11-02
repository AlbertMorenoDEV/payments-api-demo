package application

import (
	"context"
	"fmt"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/wallet/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/application"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"github.com/AlbertMorenoDEV/payments-api-demo/pkg/command"
	"reflect"
)

type CreateWalletCommand struct {
	UserId   string `json:"user_id"`
	Currency string `json:"currency"`
}

func (c CreateWalletCommand) CommandName() string {
	return "create_wallet_command"
}

type CreateWalletCommandHandler struct {
	repository           domain.WalletRepository
	timeProvider         application.TimeProvider
	domainEventPublisher sharedDomain.DomainEventPublisher
}

func NewCreateWalletCommandHandler(
	repository domain.WalletRepository,
	timeProvider application.TimeProvider,
	domainEventPublisher sharedDomain.DomainEventPublisher,
) *CreateWalletCommandHandler {
	return &CreateWalletCommandHandler{
		repository:           repository,
		timeProvider:         timeProvider,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h CreateWalletCommandHandler) Handle(_ context.Context, command command.Command) error {
	c, ok := command.(CreateWalletCommand)
	if !ok {
		return fmt.Errorf("invalid command, expected <CreateWalletCommand> and got <%s>", reflect.TypeOf(command))
	}

	userId, err := sharedUserId.New(c.UserId)
	if err != nil {
		return err
	}

	transaction := domain.NewWallet(
		*userId,
		money.Currency(c.Currency),
		h.timeProvider.Now(),
	)

	events := transaction.PullDomainEvents()

	err = h.repository.Save(transaction)

	if err != nil {
		return err
	}

	go h.domainEventPublisher.Publish(events)

	return err
}
