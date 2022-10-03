package domain

type TransactionRepository interface {
	Save(transaction *Transaction) error
}
