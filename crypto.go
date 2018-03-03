package main

import (
  "log"
  "github.com/levigross/grequests"
)

func GetCryptoRates() map[string]float64 {
  log.Println("GetCryptoRates")

  url := "https://min-api.cryptocompare.com/data/price"
  requestOptions := &grequests.RequestOptions{Params: map[string]string{"fsym": "USD", "tsyms": "BTC,ETH,LTE,BCH"}}

  resp, err := grequests.Get(url, requestOptions)

  log.Println(resp.String())

  currencyRates := map[string]float64{}

  err = resp.JSON(currencyRates)
  if err != nil {
    log.Fatalln("Json decode error")
  }

  return currencyRates
}
