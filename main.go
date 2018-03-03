package main

func main() {
	// SendSMS("+420725543082", "ty pidor")
	// GetCryptoRates()
	SaveUser(userInfo{
		phone: "+420725543082",
		senderName: "Karamba",
		currency: "BTC",
		percentage: "10"})
}
