package models

import (
	"time"
)

type Wallet struct {
	ID              int
	UID             string
	UserID          string
	Token           string
	CallbackURL     string
	CurrentBalance  float64
	MobileNumber    string
	NetworkOperator string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Transactions    []Transaction `gorm:"ForeignKey:WalletID"`
}
