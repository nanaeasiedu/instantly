package payloads

type Settings struct {
	Token           string `form:"token"`
	CallbackURL     string `form:"callbackUrl"`
	NetworkOperator string `form:"networkOperator"`
	MobileNumber    string `form:"mobileNumber"`
}
