package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/SaidovZohid/market-project/storage/repo"
	"github.com/SaidovZohid/market-project/storage/postgres"
)

type StorageI interface {
	Debt() repo.DebtStorageI
}

type StoragePg struct {
	debtRepo repo.DebtStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &StoragePg{
		debtRepo: postgres.NewDebt(db),
	}
}

func (s *StoragePg) Debt() repo.DebtStorageI {
	return s.debtRepo
}
