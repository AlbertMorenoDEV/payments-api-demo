package persistence

import (
	"database/sql"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/balance/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	pkgMysql "github.com/AlbertMorenoDEV/payments-api-demo/pkg/mysql"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-sql-driver/mysql"
)

const balanceMysqlTable = "users_balance"

type MysqlBalanceRepository struct {
	mysqlClient *sql.DB
}

func NewMysqlBalanceRepository(mysqlClient *sql.DB) *MysqlBalanceRepository {
	return &MysqlBalanceRepository{mysqlClient: mysqlClient}
}

func (r MysqlBalanceRepository) Find(id sharedUserId.UserId) (*domain.Balance, error) {
	queryBuilder := sq.Select("user_id", "amount", "currency").
		From(balanceMysqlTable).
		Where(sq.Eq{"user_id": id.Value()})

	res, err := queryBuilder.RunWith(r.mysqlClient).Query()

	if err != nil {
		return nil, err
	}

	defer pkgMysql.CloseRows(res)

	if !res.Next() {
		return domain.NewBalance(id, money.NewSGD(money.NewZeroAmount())), nil
	}

	var (
		userId   string
		amount   int64
		currency string
	)

	err = res.Scan(&userId, &amount, &currency)

	if err != nil {
		return nil, err
	}

	return domain.NewBalanceFromPrimitives(userId, amount, currency)
}

func (r MysqlBalanceRepository) Save(balance *domain.Balance) error {
	rows, err := sq.Replace(balanceMysqlTable).Columns(
		"user_id",
		"amount",
		"currency",
	).Values(
		balance.UserId().Value(),
		balance.Amount().Amount().Int64(),
		balance.Amount().Currency().String(),
	).RunWith(r.mysqlClient).Query()

	defer pkgMysql.CloseRows(rows)

	_, ok := err.(*mysql.MySQLError)
	if !ok {
		return err
	}

	return err
}
