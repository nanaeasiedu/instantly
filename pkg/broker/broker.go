package broker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ngenerio/instantly/pkg/payments"

	"errors"

	"strings"

	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
)

type Broker struct {
	ClientId     string
	ClientSecret string
	Token        string
	BaseURL      string
	BrokerSender string
}

func NewBroker(clientId, clientSecret, brokerToken, brokerSender, baseUrl string) *Broker {
	broker := &Broker{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Token:        brokerToken,
		BaseURL:      baseUrl,
		BrokerSender: brokerSender,
	}

	return broker
}

func (b *Broker) DebitWallet(requestDetails payments.MPaymentRequest) payments.MPaymentResponse {
	request := new(Request)
	request.Amount = requestDetails.GetAmount()
	request.Provider = requestDetails.GetNetwork()
	request.MSISDN = requestDetails.GetNumber()
	request.ReceiverName = requestDetails.GetName()
	request.Sender = b.BrokerSender
	request.Type = "withdraw"
	request.Token = b.Token

	if request.Provider == "VODAFONE" {
		request.ReceiveToken = requestDetails.GetReceiveToken()
	}

	return b.NewRequest(request)
}

func (b *Broker) CreditWallet(requestDetails payments.MPaymentRequest) payments.MPaymentResponse {
	request := new(Request)
	request.Amount = requestDetails.GetAmount()
	request.Provider = requestDetails.GetNetwork()
	request.MSISDN = requestDetails.GetNumber()
	request.ReceiverName = requestDetails.GetName()
	request.Sender = b.BrokerSender
	request.Type = "deposit"
	request.Token = b.Token

	if request.Provider == "VODAFONE" {
		request.ReceiveToken = requestDetails.GetReceiveToken()
	}

	return b.NewRequest(request)
}

func (b *Broker) NewRequest(data interface{}) payments.MPaymentResponse {
	newRequest := gorequest.New().SetBasicAuth(b.ClientId, b.ClientSecret)

	resp, body, errs := newRequest.
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Post(b.BaseURL + "/mobilemoney").
		Send(data).
		End()

	if errs != nil {
		return payments.NewResponse(nil, b.NewError(errs))
	}

	body = strings.Replace(body, "\x00", "", -1)
	response := Response{}
	withoutNullBytes := bytes.Trim([]byte(body), "\x00")
	body = string(withoutNullBytes[:])

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			log.Warn(fmt.Sprintf("%v %s", err, "JSON syntax error"))
		case *json.InvalidUnmarshalError:
			log.Warn(fmt.Sprintf("%v %s", err, "Invalid unmarshal error"))
		case *json.InvalidUTF8Error:
			log.Warn(fmt.Sprintf("%v %s", err, "UTF-8 style error"))
		case *json.UnmarshalFieldError:
			log.Warn(fmt.Sprintf("%v %s", err, "Unmarshal field error"))
		}

		return payments.NewResponse(nil, err)
	}

	if resp.StatusCode != http.StatusOK {
		return payments.NewResponse(nil, errors.New(response.Description))
	}

	return payments.NewResponse(response, nil)
}

func (b *Broker) NewError(err interface{}) error {
	return fmt.Errorf("Error occured executing request %q", err)
}
