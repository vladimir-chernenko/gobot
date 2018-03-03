package main

import (
  "github.com/go-redis/redis"
  "encoding/json"
  "log"
)

var client = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0})


func SaveUser(person userInfo) {
  currencyRates := GetCryptoRates()
  person.ActualPrice = currencyRates[person.Currency]

  b, err := json.Marshal(person)

  if err != nil {
      log.Println(err)
      return
  }

  err = client.Set(person.Phone, b, 0).Err()
  if err != nil {
    panic(err)
  }
}


func UpdateCurrencyPrice() {
  allPhoneNumbers, err := client.Keys("*").Result()

  log.Printf("All phone numbers %s", allPhoneNumbers)

  if err != nil {
    log.Fatalln("Can not get all numbers from Redis")
  }

  currencyRates := GetCryptoRates()

  for _, phone := range allPhoneNumbers {
    personJson, _ := client.Get(phone).Result()
    person := userInfo{}

    _ = json.Unmarshal([]byte(personJson), &person)

    person.ActualPrice = currencyRates[person.Currency]

    b, _ := json.Marshal(person)

    _ = client.Set(person.Phone, b, 0).Err()
  }

  log.Println(allPhoneNumbers)
}
