package domain

import (
	sharedUserId "github.com/AlbertMorenoDEV/payments-api-demo/internal/pkg/domain/user-id"
)

type WalletRepository interface {
	Find(userId sharedUserId.UserId) (*Wallet, error)
	Save(wallet *Wallet) error
}
