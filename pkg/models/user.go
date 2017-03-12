package models

import (
	"time"
)

type User struct {
	ID           string
	EmailAddress string
	Password     string
	PasswordHash string
	Wallets      []Wallet `gorm:"ForeignKey:UserID`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
