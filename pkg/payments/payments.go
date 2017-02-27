package payments

import (
	"github.com/ngenerio/instantly/pkg/utils"
	"github.com/satori/go.uuid"
)

type MPayment interface {
	CreditWallet(request MPaymentRequest) MPaymentResponse
	DebitWallet(request MPaymentRequest) MPaymentResponse
}

type ProviderResponse interface {
	GetProviderID() string
	GetResponse() interface{}
	GetTransactionID() string
}

type MPaymentRequest interface {
	GetName() string
	GetNumber() string
	GetAmount() float64
	GetNetwork() string
	GetReferenceID() string
	GetReceiveToken() string
	GetType() string
}

type MPaymentResponse interface {
	IsError() bool
	Error() string
	GetResponseData() interface{}
	GetTransactionID() string
	GetNetworkID() string
}

type Request struct {
	Name         string  `json:"name,omitempty"`
	MobileNumber string  `json:"phoneNumber"`
	Amount       float64 `json:"amount,omitempty"`
	MNO          string  `json:"mno,omitempty"`
	reference    string
	receiveToken string
	Type         string `json:"type"`
}

type TransferRequest struct {
	*Request
	From string `json:"from"`
	To   string `json:"to"`
}

func NewTransferRequest() *TransferRequest {
	newTransfer := new(TransferRequest)
	newTransfer.Request = NewReqeust()

	return newTransfer
}

func NewReqeust() *Request {
	newRequest := new(Request)
	newRequest.reference = uuid.NewV4().String()
	newRequest.receiveToken = utils.GenerateRandomString(5)

	return newRequest
}

func (m *Request) GetName() string { return m.Name }

func (m *Request) GetNumber() string { return m.MobileNumber }

func (m *Request) GetAmount() float64 { return m.Amount }

func (m *Request) GetNetwork() string { return m.MNO }

func (m *Request) GetReferenceID() string { return m.reference }

func (m *Request) GetReceiveToken() string { return m.receiveToken }

func (m *Request) GetType() string { return m.Type }

type Response struct {
	provider ProviderResponse
	Status   string
	Message  string
	err      error
}

func NewResponse(data ProviderResponse, err error) *Response {
	var response *Response = new(Response)
	response.provider = data
	if err != nil {
		response.err = err
		response.Message = err.Error()
		response.Status = "error"

		return response
	}

	response.Status = "success"

	return response
}

func (response *Response) IsError() bool { return response.err != nil }

func (response *Response) Error() string { return response.err.Error() }

func (response *Response) GetResponseData() interface{} {
	return response.provider.GetResponse()
}

func (response *Response) GetTransactionID() string {
	return response.provider.GetTransactionID()
}

func (response *Response) GetNetworkID() string {
	return response.provider.GetProviderID()
}
