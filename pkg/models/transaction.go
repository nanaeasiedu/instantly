package models

import (
	"time"

	"github.com/ngenerio/instantly/pkg/payments"
	"github.com/ngenerio/instantly/pkg/utils"
)

const (
	StatusPending = "PENDING"
	StatusFailed  = "FAILED"
	StatusSuccess = "SUCCESS"
)

type Transaction struct {
	Id           int       `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CompletedAt  time.Time `json:"completedAt"`
	Amount       float64   `json:"amount"`
	Type         string    `json:"type"`
	MNO          string    `json:"mno"`
	Reference    string    `json:"reference"`
	Message      string    `json:"message"`
	MobileNumber string    `json:"mobileNumber"`
	ReceiveToken string    `json:"receiveToken"`
	NetworkID    string    `json:"networkID"`
	Status       string    `json:"status"`
	ReferenceID  string    `json:"referenceID"`
	UserID       int       `json:"userID"`
}

func CreateTransaction(paymentRequest payments.MPaymentRequest, typeOfTrx string, user *User) (*Transaction, error) {
	trxDataStore := new(Transaction)
	trxDataStore.Amount = paymentRequest.GetAmount()
	trxDataStore.MNO = paymentRequest.GetNetwork()
	trxDataStore.MobileNumber = paymentRequest.GetNumber()
	trxDataStore.ReferenceID = paymentRequest.GetReferenceID()
	trxDataStore.Status = StatusPending
	trxDataStore.CreatedAt = time.Now()
	trxDataStore.Type = typeOfTrx
	trxDataStore.ReceiveToken = paymentRequest.GetReceiveToken()
	trxDataStore.UserID = user.ID

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

func (trx *Transaction) GetTransaction(queryParam map[string]interface{}) error {
	err := db.Where(queryParam).First(trx).Error
	return err
}

func (trx *Transaction) Validate() error {
	phoneNumber, err := utils.ParsePhoneNumber(trx.MobileNumber)

	if err != nil {
		return err
	}

	trx.MobileNumber = phoneNumber
	return nil
}
