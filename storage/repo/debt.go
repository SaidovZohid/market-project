package repo

import "time"

type Debt struct {
	ID                    int64
	FirstName             string
	LastName              string
	PhoneNumber           string
	AdditionalPhoneNumber *string
	AddressWork           string
	SellerFullName        string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             time.Time
}

type DebtStorageI interface {
	Create(debt *Debt) (*Debt, error)
	Get(debt_id int64) (*Debt, error)
	Update(debt *Debt) error
	Delete(debt_id int64) error
	GetAll(params *GetAllParams) (*GetAllDebts, error)
}

type GetAllParams struct {
	Limit           int64
	Page            int64
	FirstName       string
	LastName        string
	PhoneNumber     string
	AdditionalPhone string
	SellerFullName  string
}

type GetAllDebts struct {
	Debts []*Debt
	Count int64
}
