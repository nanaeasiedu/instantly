package api

import (
	"net/http"
	"reflect"

	"github.com/ngenerio/instantly/pkg/broker"
	"github.com/ngenerio/instantly/pkg/models"
	"github.com/ngenerio/instantly/pkg/payments"

	"github.com/labstack/echo"
	"github.com/ngenerio/instantly/pkg/config"
	log "github.com/sirupsen/logrus"
)

const (
	Debit  = "debit"
	Credit = "credit"
)

var paymentsSolution payments.MPayment = broker.NewBroker(config.Settings.UnityClientID, config.Settings.UnityClientSecret, config.Settings.BrokerToken, config.Settings.BrokerSender, config.Settings.BrokerBaseURL)

func HandlePayments(c echo.Context) error {
	var request payments.MPaymentRequest = payments.NewReqeust()
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: "error", Message: err.Error()})
	}

	log.Info("Log the api request body", request)
	requestInterface := reflect.ValueOf(request).Interface()
	paymentRequest := requestInterface.(payments.MPaymentRequest)

	newTransaction, err := models.CreateTransaction(paymentRequest, request.GetType())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: "error", Message: err.Error()})
	}

	var response payments.MPaymentResponse
	if request.GetType() == Debit {
		response = paymentsSolution.DebitWallet(paymentRequest)
	} else {
		response = paymentsSolution.CreditWallet(paymentRequest)
	}

	if response.IsError() {
		newTransaction.Message = response.Error()
		newTransaction.Reference = "N/A"
		newTransaction.Status = models.StatusFailed
		_ = newTransaction.Update()
		return c.JSON(http.StatusInternalServerError, Response{Status: "error", Message: response.Error()})
	}

	responseToSend := &Response{}
	responseToSend.Status = "success"
	newTransaction.Reference = response.GetTransactionID()
	newTransaction.NetworkID = response.GetNetworkID()

	if request.GetType() == Debit {
		newTransaction.Status = models.StatusPending
	} else {
		newTransaction.Status = models.StatusSuccess
	}

	err = newTransaction.Update()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: "error", Message: err.Error()})
	}

	responseToSend.Data = newTransaction
	return c.JSON(http.StatusOK, responseToSend)
}

func HandlePaymentsTransfer(e echo.Context) error {
	return nil
}

func HandleCallback(e echo.Context) error {
	return nil
}
