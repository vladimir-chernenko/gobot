package main

import (
  "github.com/go-redis/redis"
  "encoding/json"
  "fmt"
)


var client = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0})


type userInfo struct {
  phone string
  senderName string
  currency string
  percentage string
  actualPrice float64
}

func SaveUser(person userInfo) {
  currencyRates := GetCryptoRates()
  person.actualPrice = currencyRates[person.currency]

  b, err := json.Marshal(person)

  if err != nil {
      fmt.Println(err)
      return
  }

  err = client.Set(person.phone, b, 0).Err()
  if err != nil {
    panic(err)
  }
}


func UpdateCurrencyPrice() {
  allPhoneNumbers = client.Keys("*")
  fmt.Println(allPhoneNumbers)
}
