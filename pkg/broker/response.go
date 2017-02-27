package broker

type Response struct {
	ID          string `json:"Id,omitempty"`
	Customer    string `json:"Customer,omitempty"`
	ForeignID   string `json:"ForeignId,omitempty"`
	ProviderID  string `json:"ProviderId,omitempty"`
	ErrorCode   string `json:"ErrorCode,omitempty"`
	Description string `json:"Description,omitempty"`
}

func (r Response) GetProviderID() string { return r.ProviderID }

func (r Response) GetResponse() interface{} { return r }

func (r Response) GetTransactionID() string { return r.ID }
