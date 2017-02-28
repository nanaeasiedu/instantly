package payments

type ProviderResponse interface {
	GetProviderID() string
	GetResponse() interface{}
	GetTransactionID() string
}

type MPaymentResponse interface {
	IsError() bool
	Error() string
	GetResponseData() interface{}
	GetTransactionID() string
	GetNetworkID() string
}

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
