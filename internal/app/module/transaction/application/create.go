package application

import (
	"context"
	"fmt"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain"
	transactionId "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain/transaction-id"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/application"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"github.com/AlbertMorenoDEV/payments-api-demo/pkg/command"
	"reflect"
)

type CreateTransactionCommand struct {
	TransactionId     string `json:"transaction_id"`
	UserId            string `json:"user_id"`
	DestinationUserId string `json:"destination_user_id"`
	Amount            int64  `json:"amount"`
	Currency          string `json:"currency"`
}

func (c CreateTransactionCommand) CommandName() string {
	return "create_transaction_command"
}

type CreateTransactionCommandHandler struct {
	repository           domain.TransactionRepository
	timeProvider         application.TimeProvider
	domainEventPublisher sharedDomain.DomainEventPublisher
}

func NewCreateTransactionCommandHandler(
	repository domain.TransactionRepository,
	timeProvider application.TimeProvider,
	domainEventPublisher sharedDomain.DomainEventPublisher,
) *CreateTransactionCommandHandler {
	return &CreateTransactionCommandHandler{
		repository:           repository,
		timeProvider:         timeProvider,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h CreateTransactionCommandHandler) Handle(_ context.Context, command command.Command) error {
	createTransactionCommand, ok := command.(CreateTransactionCommand)
	if !ok {
		return fmt.Errorf("invalid command, expected <CreateTransactionCommand> and got <%s>", reflect.TypeOf(command))
	}

	transId, err := transactionId.New(createTransactionCommand.TransactionId)
	if err != nil {
		return err
	}

	userId, err := sharedUserId.New(createTransactionCommand.UserId)
	if err != nil {
		return err
	}

	destinationUserId, err := sharedUserId.New(createTransactionCommand.DestinationUserId)
	if err != nil {
		return err
	}

	transaction := domain.NewTransaction(
		*transId,
		*userId,
		*destinationUserId,
		money.NewFromPrimitives(createTransactionCommand.Amount, createTransactionCommand.Currency),
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
