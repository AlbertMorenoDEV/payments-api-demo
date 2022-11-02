package domain

import (
	"errors"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain"
	transactionId "github.com/AlbertMorenoDEV/payments-api-demo/internal/app/module/transaction/domain/transaction-id"
	sharedDomain "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain"
	"github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/money"
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
	"github.com/AlbertMorenoDEV/payments-api-demo/pkg/event"
	"sync"
	"time"
)

type Wallet struct {
	userId           sharedUserId.UserId
	balance          money.Money
	blockedBalance   money.Money
	creationDate     time.Time
	modificationDate time.Time

	transactions map[transactionId.TransactionId]*Transaction

	mutex         sync.Mutex
	eventRecorder *sharedDomain.EventRecorder

	version int64
}

func NewWallet(userId sharedUserId.UserId, currency money.Currency, now time.Time) *Wallet {
	w := &Wallet{
		eventRecorder: sharedDomain.InitEventRecorder(),
	}

	e := NewWalletCreated(userId, currency, now)
	w.eventRecorder.Record(e)
	w.applyEvent(e)

	return w
}

func (w *Wallet) applyEvent(e event.Event) {
	switch e.(type) {
	case *WalletCreated:
		w.applyWalletCreatedEvent(e.(*WalletCreated))
	default:
		panic(errors.New("unhandled domain event"))
	}
}

func (w *Wallet) applyWalletCreatedEvent(e *WalletCreated) {
	w.userId = e.userId
	w.balance = money.New(money.NewZeroAmount(), e.Currency())
	w.blockedBalance = money.New(money.NewZeroAmount(), e.Currency())
	w.creationDate = e.CreatedAt()
	w.modificationDate = e.CreatedAt()
	w.version++
}

func (w *Wallet) StartTransfer(
	transactionId transactionId.TransactionId,
	destinationUserId sharedUserId.UserId,
	amount money.Money,
	now time.Time,
) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if _, ok := w.transactions[transactionId]; ok {
		return domain.NewTransactionAlreadyExist(transactionId)
	}

	if w.balance.IsLessThan(amount) {
		return errors.New("not enough funds")
	}

	t := &Transaction{
		transactionId:     transactionId,
		userId:            w.userId,
		originUserId:      w.userId,
		destinationUserId: destinationUserId,
		amount:            amount,
		creationDate:      now,
		status:            transactionStatusPending,
	}

	w.transactions[transactionId] = t

	w.eventRecorder.Record(NewTransactionCreated(*t))

	w.balance = w.balance.Subtract(amount)
	w.blockedBalance = w.blockedBalance.Add(amount)
	w.modificationDate = now

	return nil
}

func (w *Wallet) AcceptTransfer(
	transactionId transactionId.TransactionId,
	originUserId sharedUserId.UserId,
	amount money.Money,
	now time.Time,
) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if _, ok := w.transactions[transactionId]; ok {
		w.eventRecorder.Record(NewTransactionRejected(transactionId))

		return domain.NewTransactionAlreadyExist(transactionId)
	}

	t := &Transaction{
		transactionId:     transactionId,
		userId:            w.userId,
		originUserId:      originUserId,
		destinationUserId: w.userId,
		amount:            amount,
		creationDate:      now,
		status:            transactionStatusAccepted,
	}

	w.transactions[transactionId] = t

	w.eventRecorder.Record(NewTransactionAccepted(*t))

	w.balance = w.balance.Add(amount)
	w.modificationDate = now

	return nil
}

func (w *Wallet) FinishTransfer(
	transactionId transactionId.TransactionId,
	amount money.Money,
	now time.Time,
) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	t, ok := w.transactions[transactionId]
	if !ok {
		w.eventRecorder.Record(NewTransactionFinalizationRejected(transactionId))

		return domain.NewTransactionNotExists(transactionId)
	}

	if err := t.verify(now); err != nil {
		return err
	}

	w.eventRecorder.Record(NewTransactionFinished(*t))

	w.blockedBalance = w.blockedBalance.Subtract(amount)
	w.balance = w.balance.Add(amount)
	w.modificationDate = now

	return nil
}

func (w *Wallet) UserId() sharedUserId.UserId {
	return w.userId
}

func (w *Wallet) Balance() money.Money {
	return w.balance
}

func (w *Wallet) BlockedBalance() money.Money {
	return w.blockedBalance
}

func (w *Wallet) CreationDate() time.Time {
	return w.creationDate
}

func (w *Wallet) ModificationDate() time.Time {
	return w.modificationDate
}

func (w *Wallet) Transactions() map[transactionId.TransactionId]*Transaction {
	return w.transactions
}

func (w *Wallet) PullDomainEvents() sharedDomain.DomainEvents {
	return w.eventRecorder.Pull()
}
