package infrastructure

import (
	"database/sql"
	balanceDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/domain"
	balanceInfrastructurePersistence "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/infrastructure/persistence"
	balanceUi "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/ui"
	transactionDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain"
	transactionInfrastructurePersistence "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/infrastructure/persistence"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/application"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/pkg/mysql"
	"go.uber.org/zap"
)

type Services struct {
	logger                *zap.Logger
	timeProvider          application.TimeProvider
	balanceRepository     balanceDomain.BalanceRepository
	transactionRepository transactionDomain.TransactionRepository
	domainEventPublisher  domain.DomainEventPublisher
	balanceProjector      *balanceUi.BalanceProjector
}

func (s Services) Logger() *zap.Logger {
	return s.logger
}

func (s Services) TimeProvider() application.TimeProvider {
	return s.timeProvider
}

func (s Services) BalanceRepository() balanceDomain.BalanceRepository {
	return s.balanceRepository
}

func (s Services) TransactionRepository() transactionDomain.TransactionRepository {
	return s.transactionRepository
}

func (s Services) DomainEventPublisher() domain.DomainEventPublisher {
	return s.domainEventPublisher
}

func (s Services) BalanceProjector() *balanceUi.BalanceProjector {
	return s.balanceProjector
}

func BuildServices(config *Config, domainEvents chan domain.DomainEvent) *Services {
	sqlClient := buildMysqlClient(config)
	logger := buildLogger()

	balanceRepository := balanceInfrastructurePersistence.NewMysqlBalanceRepository(sqlClient)

	return &Services{
		logger:                logger,
		timeProvider:          SystemTimeProvider{},
		balanceRepository:     balanceRepository,
		transactionRepository: transactionInfrastructurePersistence.NewMysqlTransactionRepository(sqlClient),
		domainEventPublisher:  NewChannelDomainEventPublisher(domainEvents),
		balanceProjector:      balanceUi.NewBalanceProjector(domainEvents, logger, balanceRepository),
	}
}

func buildMysqlClient(config *Config) *sql.DB {
	sqlClient, err := mysql.Connect(
		config.mysqlUser,
		config.mysqlPass,
		config.mysqlHost,
		config.mysqlPort,
		config.mysqlDb,
	)
	if err != nil {
		panic(err)
	}

	return sqlClient
}

func buildLogger() *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return l
}
