package ui

import (
	"context"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/application"
	domain "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/domain"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"go.uber.org/zap"
)

type BalanceProjector struct {
	domainEvents chan sharedDomain.DomainEvent
	logger       *zap.Logger
	repository   domain.BalanceRepository
}

func NewBalanceProjector(
	domainEvents chan sharedDomain.DomainEvent,
	logger *zap.Logger,
	repository domain.BalanceRepository,
) *BalanceProjector {
	return &BalanceProjector{
		domainEvents: domainEvents,
		logger:       logger,
		repository:   repository,
	}
}

func (p *BalanceProjector) Start(ctx context.Context) {
	ch := application.NewUpdateBalanceCommandHandler(p.repository)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			for event := range p.domainEvents {
				p.logger.Info("New event! " + event.Name())

				err := ch.Handle(event)
				if err != nil {
					p.logger.Error("Error on Balance projector", zap.Error(err))
				}
			}
		}
	}
}
