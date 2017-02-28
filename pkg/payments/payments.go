package payments

type MPayment interface {
	CreditWallet(request MPaymentRequest) MPaymentResponse
	DebitWallet(request MPaymentRequest) MPaymentResponse
}
