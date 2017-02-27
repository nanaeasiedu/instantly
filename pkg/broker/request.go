package broker

type Request struct {
	MSISDN       string  `json:"receiverPhone"`
	ReceiverName string  `json:"receiverName"`
	Amount       float64 `json:"amount"`
	ForeignID    string  `json:"foreignId,omitempty"`
	Provider     string  `json:"provider"`
	Sender       string  `json:"sender"`
	Token        string  `json:"token"`
	ReceiveToken string  `json:"receiveToken,omitempty"`
	Type         string  `json:"type"`
	CallbackURL  string  `json:"callbackUrl"`
}
