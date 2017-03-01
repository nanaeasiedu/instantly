package api

import (
	"net/http"
	"time"

	"github.com/ngenerio/instantly/pkg/broker"
	"github.com/ngenerio/instantly/pkg/models"
	"github.com/ngenerio/instantly/pkg/payments"

	"strings"

	"github.com/labstack/echo"
	"github.com/ngenerio/instantly/pkg/config"
	log "github.com/sirupsen/logrus"
)

const (
	Debit  = "debit"
	Credit = "credit"
)

var paymentsSolution payments.MPayment = broker.NewBroker(config.Settings.UnityClientID, config.Settings.UnityClientSecret, config.Settings.BrokerToken, config.Settings.BrokerSender, config.Settings.BrokerBaseURL, config.Settings.BrokerCallbackURL)

func HandlePayments(c echo.Context) error {
	request := payments.NewReqeust()
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: "error", Message: err.Error()})
	}

	log.Info("Log the api request body", request)
	newTransaction, err := models.CreateTransaction(request, request.GetType())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: "error", Message: err.Error()})
	}

	var response payments.MPaymentResponse
	if request.GetType() == Debit {
		response = paymentsSolution.DebitWallet(request)
	} else {
		response = paymentsSolution.CreditWallet(request)
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

	newTransaction.Status = models.StatusPending
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

func HandleCallback(c echo.Context) error {
	transactionId := c.QueryParam("transactionId")
	statusOfTrx := c.QueryParam("status")

	log.Info("Callback url called by broker with params: ", transactionId, statusOfTrx)
	newTrx := new(models.Transaction)
	err := newTrx.GetTransaction(map[string]interface{}{"reference": transactionId})

	if err != nil {
		log.Error("Error occured retreiving transaction from db", err)
		return nil
	}

	if strings.ToLower(statusOfTrx) == "failed" {
		newTrx.Message = "Payment failed"
		newTrx.Status = models.StatusFailed
	} else {
		newTrx.Message = "Payment was successful"
		newTrx.Status = models.StatusSuccess
	}

	newTrx.CompletedAt = time.Now()

	newTrx.Update()
	return nil
}
