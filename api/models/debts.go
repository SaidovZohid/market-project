package models

import "time"

type CreateDebt struct {
	FirstName             string  `json:"first_name" binding:"required,min=2,max=50"`
	LastName              string  `json:"last_name" binding:"required,min=2,max=50"`
	PhoneNumber           string  `json:"phone_number" binding:"required,min=9,max=30"`
	AdditionalPhoneNumber *string `json:"additional_phone_number" binding:"required,min=9,max=30"`
	AddressWork           string  `json:"address_work" binding:"required"`
	SellerFullName        string  `json:"seller_fullname" binding:"required,min=4,max=50"`
}

type GetDebt struct {
	ID                    int64     `json:"id"`
	FirstName             string    `json:"first_name"`
	LastName              string    `json:"last_name" `
	PhoneNumber           string    `json:"phone_number"`
	AdditionalPhoneNumber *string   `json:"additional_phone_number"`
	AddressWork           string    `json:"address_work"`
	SellerFullName        string    `json:"seller_fullname"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	DeletedAt             time.Time `json:"deleted_at"`
}

type GetAfterCreate struct {
	ID                    int64     `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
}

type GetDebts struct {
	Debts []*GetDebt `json:"debts"`
	Count int64      `json:"counts"`
}

type GetAllParams struct {
	Limit           int64  `json:"limit"`
	Page            int64  `json:"page"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	AdditionalPhone string `json:"additional_phone_number"`
	SellerFullName  string `json:"seller_fullname"`
}
