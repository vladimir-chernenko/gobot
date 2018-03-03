package main

import (
	"time"
  "log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)


type userInfo struct {
  Phone string `json:"Phone"`
  Currency string `json:"Currency"`
  Percentage string `json:"Percentage"`
  ActualPrice float64 `json:"ActualPrice"`
}


func main() {
	b, err := tb.NewBot(tb.Settings{
	Token: os.Getenv("TELEGRAM_API_TOKEN"),
	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	user := &userInfo{}

	if err != nil {
		log.Fatal(err)
		return
	}

	handleStart(b, user)

	b.Start()
	// SendSMS("+420725543082", "ty pidor")
	// GetCryptoRates()
// 	SaveUser(userInfo{
// 		Phone: "+420725543082",
// 		Currency: "BTC",
// 		Percentage: "10"})
// 	UpdateCurrencyPrice()
}
