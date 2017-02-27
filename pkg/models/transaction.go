package models

import (
	"errors"
	"time"

	"github.com/ngenerio/instantly/pkg/payments"
	"github.com/ttacon/libphonenumber"
)

const (
	StatusPending = "PENDING"
	StatusFailed  = "FAILED"
	StatusSuccess = "SUCCESS"
)

var errorInvalidPhoneNumber = errors.New("Invalid phone number. Ensure number is in international format (23327xxxxxxx)")

type Transaction struct {
	Id           int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Amount       float64
	Type         string
	MNO          string
	Reference    string
	Message      string
	MobileNumber string
	ReceiveToken string
	NetworkID    string `gorm:"column:network_id"`
	Status       string
	ReferenceID  string `gorm:"column:reference_id"`
}

func CreateTransaction(paymentRequest payments.MPaymentRequest, typeOfTrx string) (*Transaction, error) {
	trxDataStore := new(Transaction)
	trxDataStore.Amount = paymentRequest.GetAmount()
	trxDataStore.MNO = paymentRequest.GetNetwork()
	trxDataStore.MobileNumber = paymentRequest.GetNumber()
	trxDataStore.ReferenceID = paymentRequest.GetReferenceID()
	trxDataStore.Status = StatusPending
	trxDataStore.CreatedAt = time.Now()
	trxDataStore.Type = typeOfTrx
	trxDataStore.ReceiveToken = paymentRequest.GetReceiveToken()

	if err := trxDataStore.Validate(); err != nil {
		return nil, err
	}

	err := db.Create(trxDataStore).Error
	return trxDataStore, err
}

func (trx *Transaction) Update() error {
	trx.UpdatedAt = time.Now()
	return db.Save(trx).Error
}

func (trx *Transaction) Validate() error {
	phoneNumber, err := libphonenumber.Parse(trx.MobileNumber, "GH")

	if err != nil {
		return err
	}

	ok := libphonenumber.IsValidNumberForRegion(phoneNumber, "GH")

	if !ok {
		return errorInvalidPhoneNumber
	}

	return nil
}
