package infrastructure

import (
	"database/sql"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/wallet/domain"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	pkgMysql "github.com/AlbertMorenoDEV/payments-api-demo/pkg/mysql"
	sq "github.com/Masterminds/squirrel"
)

const walletMysqlTable = "wallet"
const transactionsMysqlTable = "wallet_transactions"

type MysqlWalletRepository struct {
	mysqlClient *sql.DB
}

func (r MysqlWalletRepository) Find(userId sharedUserId.UserId) (*domain.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (r MysqlWalletRepository) Save(wallet *domain.Wallet) error {
	tx, err := r.mysqlClient.Begin()
	if err != nil {
		return err
	}

	rows, err := sq.Replace(walletMysqlTable).Columns(
		"user_id",
		"balance_amount",
		"blocked_amount",
		"currency",
		"created_date",
		"updated_date",
	).Values(
		wallet.UserId().Value(),
		wallet.Balance().Amount().Int64(),
		wallet.BlockedBalance().Amount().Int64(),
		wallet.Balance().Currency().String(),
		wallet.CreationDate(),
		wallet.ModificationDate(),
	).RunWith(r.mysqlClient).Query()

	defer pkgMysql.CloseRows(rows)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	for _, transaction := range wallet.Transactions() {
		err := r.saveTransaction(transaction)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}

	return tx.Commit()
}

func (r MysqlWalletRepository) saveTransaction(transaction *domain.Transaction) error {
	rows, err := sq.Replace(transactionsMysqlTable).Columns(
		"transaction_id",
		"user_id",
		"origin_user_id",
		"destination_user_id",
		"amount",
		"currency",
		"creation_date",
		"modification_date",
		"status",
	).Values(
		transaction.TransactionId().Value(),
		transaction.UserId().Value(),
		transaction.OriginUserId().Value(),
		transaction.DestinationUserId().Value(),
		transaction.Amount().Amount().Float64(),
		transaction.Amount().Currency().String(),
		transaction.CreationDate(),
		transaction.ModificationDate(),
	).RunWith(r.mysqlClient).Query()

	defer pkgMysql.CloseRows(rows)

	return err
}
