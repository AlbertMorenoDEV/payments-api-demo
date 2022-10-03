package application_test

import (
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/application"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain"
	transactionId "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain/transaction-id"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"github.com/AlbertMorenoDEV/payments-api-demo/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"
)

func TestCreateTransactionSuccessfully(t *testing.T) {
	transactionRepository := new(TransactionRepositoryMock)
	timeProvider := test.NewFixedDateTimeProvider()
	domainEventPublisher := new(DomainEventPublisherMock)

	handler := application.NewCreateTransactionCommandHandler(
		transactionRepository,
		timeProvider,
		domainEventPublisher,
	)
	transactionId, _ := transactionId.New("cb41da86-d70a-4fba-8581-ecad7e06854a")
	userId, _ := sharedUserId.New("1e8c8912-7caf-43c1-99c7-267082915291")
	destinationUserId, _ := sharedUserId.New("65739819-9eb7-4593-a53e-46af7eb10b23")
	amount := money.NewFromPrimitives(4500, "USD")
	command := application.CreateTransactionCommand{
		TransactionId:     transactionId.Value(),
		UserId:            userId.Value(),
		DestinationUserId: destinationUserId.Value(),
		Amount:            amount.Amount().Int64(),
		Currency:          amount.Currency().String(),
	}
	transaction := domain.NewTransaction(
		*transactionId,
		*userId,
		*destinationUserId,
		amount,
		timeProvider.Now(),
	)
	transaction.PullDomainEvents()

	transactionRepository.ShouldSaveTransaction(transaction)

	err := handler.Handle(command)

	assert.NoError(t, err)
	domainEventPublisher.Wg().Wait()
	mock.AssertExpectationsForObjects(t, transactionRepository, domainEventPublisher)
}

type TransactionRepositoryMock struct {
	mock.Mock
}

func (r *TransactionRepositoryMock) Save(transaction *domain.Transaction) error {
	args := r.Called(transaction)

	return args.Error(0)
}

func (r *TransactionRepositoryMock) ShouldSaveTransaction(transaction *domain.Transaction) {
	r.
		On("Save", transaction).
		Once().
		Return(nil)
}

type DomainEventPublisherMock struct {
	mock.Mock
	wg sync.WaitGroup
}

func (p *DomainEventPublisherMock) Publish(events sharedDomain.DomainEvents) {
	p.Called(events)
}

func (p *DomainEventPublisherMock) Wg() *sync.WaitGroup {
	return &p.wg
}
