package persistence

import (
	"database/sql"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain"
	pkgMysql "github.com/AlbertMorenoDEV/payments-api-demo/pkg/mysql"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-sql-driver/mysql"
)

const transactionMysqlTable = "transactions"

type MysqlTransactionRepository struct {
	mysqlClient *sql.DB
}

func NewMysqlTransactionRepository(mysqlClient *sql.DB) *MysqlTransactionRepository {
	return &MysqlTransactionRepository{mysqlClient: mysqlClient}
}

func (r MysqlTransactionRepository) Save(transaction *domain.Transaction) error {
	rows, err := sq.Insert(transactionMysqlTable).Columns(
		"transaction_id",
		"user_id",
		"destination_user_id",
		"amount",
		"currency",
		"creation_date",
	).Values(
		transaction.TransactionId().Value(),
		transaction.UserId().Value(),
		transaction.DestinationUserId().Value(),
		transaction.Amount().Amount().Float64(),
		transaction.Amount().Currency().String(),
		transaction.CreationDate(),
	).RunWith(r.mysqlClient).Query()

	defer pkgMysql.CloseRows(rows)

	me, ok := err.(*mysql.MySQLError)
	if !ok {
		return err
	}

	if me.Number == 1062 {
		return domain.NewTransactionAlreadyExist(transaction.TransactionId())
	}

	return err
}
