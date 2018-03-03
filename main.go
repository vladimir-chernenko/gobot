package main


type userInfo struct {
  Phone string `json:"Phone"`
  Currency string `json:"Currency"`
  Percentage string `json:"Percentage"`
  ActualPrice float64 `json:"ActualPrice"`
}


func main() {
	// SendSMS("+420725543082", "ty pidor")
	// GetCryptoRates()
	SaveUser(userInfo{
		Phone: "+420725543082",
		Currency: "BTC",
		Percentage: "10"})
	UpdateCurrencyPrice()
}
