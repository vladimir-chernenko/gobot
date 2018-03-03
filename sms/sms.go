package sms

import (
  "github.com/levigross/grequests"
  "os"
  "log"
  "fmt"
)

func SendSMS(phone string, message string) {
  url := ""

  DEBUG := os.Getenv("GOBOT_DEBUG") == "1"

  fmt.Printf("DEBUG is %t \n", DEBUG)

  fmt.Printf("Send SMS to '%s' with text '%s' \n", phone, message)

  if DEBUG {
    url = "https://httpbin.org/post"
  } else {
    url = "https://textbelt.com/text"
  }

  textbeltKey := os.Getenv("TEXTBELT_KEY")

  data := &grequests.RequestOptions{Data: map[string]string{"phone": phone, "message": message, "key": textbeltKey}}

  // {"success":true, "textId":"10641520084164259", "quotaRemaining":29}
  // {"success":false, "error":"Invalid phone number or bad request. If your phone number contains a +, please check that you are URL encoding it properly."}
  // {"success":false,"error":"Out of quota","quotaRemaining":0}
  resp, err := grequests.Post(url, data)

  log.Println(resp.String())

  if resp.Ok != true {
    log.Println("Request did not return OK")
  }

  if err != nil {
    log.Fatalln("Unable to make request: ", err)
  }
}
