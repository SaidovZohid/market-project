package postgres

import (
	"time"

	"github.com/SaidovZohid/market-project/storage/repo"
	"github.com/jmoiron/sqlx"
)

type debtRepo struct {
	db *sqlx.DB
}

func NewDebt(db *sqlx.DB) repo.DebtStorageI {
	return &debtRepo{
		db: db,
	}
}

func (dr *debtRepo) Create(debt *repo.Debt) (*repo.Debt, error) {
	tr, err := dr.db.Begin()
	if err != nil {
		return nil, err
	}
	query := `
		INSERT INTO debts (
			first_name,
			last_name,
			phone_number,
			additional_phone_number,
			address_work,
			seller_fullname
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	var d repo.Debt
	err = tr.QueryRow(
		query,
		debt.FirstName,
		debt.LastName,
		debt.PhoneNumber,
		debt.AdditionalPhoneNumber,
		debt.AddressWork,
		debt.SellerFullName,
	).Scan(
		&d.ID,
		&d.CreatedAt,
	)
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	err = tr.Commit()
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	return &d, nil
}

func (dr *debtRepo) Get(debt_id int64) (*repo.Debt, error) {
	var debt repo.Debt

	query := `
		SELECT
			id,
			first_name,
			last_name,
			phone_number,
			additional_phone_number,
			address_work,
			seller_fullname,
			created_at,
			updated_at,
			deleted_at
		FROM debts WHERE id = $1
	`
	err := dr.db.QueryRow(
		query,
		debt_id,
	).Scan(
		&debt.ID,
		&debt.FirstName,
		&debt.LastName,
		&debt.PhoneNumber,
		&debt.AdditionalPhoneNumber,
		&debt.AddressWork,
		&debt.SellerFullName,
		&debt.CreatedAt,
		&debt.UpdatedAt,
		&debt.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &debt, nil
}

func (dr *debtRepo) Update(debt *repo.Debt) error {
	tr, err := dr.db.Begin()
	if err != nil {
		return err
	}

	query := `
		UPDATE debts SET 
			first_name = $1,
			last_name = $2,
			phone_number = $3,
			additional_phone_number = $4,
			address_work = $5,
			seller_fullname = $6,
			updated_at = $7,
		WHERE id = $8
	`
	res, err := tr.Exec(
		query,
		debt.FirstName,
		debt.LastName,
		debt.PhoneNumber,
		debt.AdditionalPhoneNumber,
		debt.AddressWork,
		debt.SellerFullName,
		time.Now(),
		debt.ID,
	)
	if err != nil {
		tr.Rollback()
		return err
	}
	result, err := res.RowsAffected()
	if err != nil {
		tr.Rollback()
		return err
	}
	if result == 0 {
		tr.Rollback()
		return err
	}

	return nil
}

func (dr *debtRepo) Delete(debt_id int64) error {
	query := `
		UPDATE debts SET deleted_at = $1 WHERE id = $2
	`
	res, err := dr.db.Exec(
		query,
		time.Now(),
		debt_id,
	)
	if err != nil {
		return err
	}
	result, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if result == 0 {
		return err
	}

	return nil
}

func (dr *debtRepo) GetAll(params *repo.GetAllParams) (*repo.GetAllDebts, error) {
	offset := (params.Page - 1) * params.Limit
	filter := " WHERE true"

	var debt repo.Debt

	query := `
		SELECT
			id,
			first_name,
			last_name,
			phone_number,
			additional_phone_number,
			address_work,
			seller_fullname,
			created_at,
			updated_at,
			deleted_at
		FROM debts WHERE id = $1
	`
	err := dr.db.QueryRow(
		query,
		debt_id,
	).Scan(
		&debt.ID,
		&debt.FirstName,
		&debt.LastName,
		&debt.PhoneNumber,
		&debt.AdditionalPhoneNumber,
		&debt.AddressWork,
		&debt.SellerFullName,
		&debt.CreatedAt,
		&debt.UpdatedAt,
		&debt.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &debt, nil
}
