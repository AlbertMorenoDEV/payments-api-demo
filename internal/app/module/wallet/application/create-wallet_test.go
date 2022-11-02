package application_test

import (
	"context"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/wallet/application"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/wallet/domain"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"github.com/AlbertMorenoDEV/payments-api-demo/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"
)

func TestCreateWalletSuccessfully(t *testing.T) {
	repository := new(WalletRepositoryMock)
	timeProvider := test.NewFixedDateTimeProvider()
	domainEventPublisher := new(DomainEventPublisherMock)

	handler := application.NewCreateWalletCommandHandler(
		repository,
		timeProvider,
		domainEventPublisher,
	)
	userId, _ := sharedUserId.New("1e8c8912-7caf-43c1-99c7-267082915291")
	currency := money.Currency("USD")
	command := application.CreateWalletCommand{
		UserId:   userId.Value(),
		Currency: string(currency),
	}
	wallet := domain.NewWallet(
		*userId,
		currency,
		timeProvider.Now(),
	)
	wallet.PullDomainEvents()
	events := sharedDomain.DomainEvents{domain.NewWalletCreated(*userId, currency, timeProvider.Now())}
	ctx := context.Background()

	repository.ShouldSaveWallet(wallet)
	domainEventPublisher.ShouldPublishDomainEvents(t, events)

	err := handler.Handle(ctx, command)

	assert.NoError(t, err)
	domainEventPublisher.Wg().Wait()
	mock.AssertExpectationsForObjects(t, repository, domainEventPublisher)
}

type WalletRepositoryMock struct {
	mock.Mock
}

func (r *WalletRepositoryMock) Find(userId sharedUserId.UserId) (*domain.Wallet, error) {
	args := r.Called(userId)

	if args[0] == nil {
		return nil, args.Error(1)
	}
	return args[0].(*domain.Wallet), nil
}

func (r *WalletRepositoryMock) Save(wallet *domain.Wallet) error {
	args := r.Called(wallet)

	return args.Error(0)
}

func (r *WalletRepositoryMock) ShouldSaveWallet(wallet *domain.Wallet) {
	r.
		On("Save", wallet).
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

func (p *DomainEventPublisherMock) ShouldPublishDomainEvents(t *testing.T, events sharedDomain.DomainEvents) {
	p.wg.Add(1)
	p.
		On("Publish", mock.MatchedBy(func(received sharedDomain.DomainEvents) bool {
			assert.Len(t, received, len(events), "Invalid number of events received")

			for i := range events {
				domainEvent := events[i].(sharedDomain.DomainEvent)
				assert.Equal(t, domainEvent.CreatedAt(), received[i].CreatedAt(), "Wrong created at value")
				assert.Equal(t, domainEvent.Name(), received[i].Name(), "Wrong name value")
				assert.Equal(t, domainEvent.AggregateId(), received[i].AggregateId(), "Wrong aggregate ID")
				assert.Equal(t, domainEvent.Data(), received[i].Data(), "Wrong data")
			}

			return true
		})).
		Once().
		Run(func(args mock.Arguments) {
			p.wg.Done()
		}).
		Return(nil)
}
