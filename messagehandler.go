package main

import (
	"log"
  "regexp"
	tb "gopkg.in/tucnak/telebot.v2"
)

func handleStart(b *tb.Bot, user *userInfo)  {
  b.Handle("/start", func(m *tb.Message) {
    b.Send(m.Sender, "Hi, I'm CryptoCollapse bot. I'm made for sending message to you friend when crypto currency he invested in collapses.")
    b.Send(m.Sender, "Please send me your friend's phone number")
    handlePhoneNumber(b, user)
  })
}

func handlePhoneNumber(b *tb.Bot, user *userInfo)  {
  b.Handle(tb.OnText, func(m *tb.Message) {
    match, _ := regexp.MatchString(`\+[0-9]+`, m.Text)
    if match {
      user.Phone = m.Text
      log.Println("Sender:", m.Sender, "Phone:" , user.Phone)
      b.Send(m.Sender, "Ok, now I need the name of crypto currency. Send me number")
			b.Send(m.Sender, "1 Bitcoin\n2 Ethereum\n3 Litecoin\n4 Bitcoin Cash")
      handleCurrency(b, user)
    } else {
      b.Send(m.Sender, "Number is not valid. Try again.")
      handlePhoneNumber(b, user)
    }
  })
}

func handleCurrency(b *tb.Bot, user *userInfo) {
  b.Handle(tb.OnText, func(m *tb.Message) {
    switch m.Text {
    case "1":
      user.Currency = "Bitcoin"
    case "2":
      user.Currency = "Ethereum"
    case "3":
      user.Currency = "Litecoin"
    case "4":
      user.Currency = "Bitcoin Cash"
    default:
      b.Send(m.Sender, "Send number from 1 to 4")
      handleCurrency(b, user)
    }
    log.Println("Sender:", m.Sender, "Currency:" , user.Currency)
    b.Send(m.Sender, "And the last thing. Send message when it drops by")
    b.Send(m.Sender, "1 25%\n2 35%\n3 50%\n4 70%")
    handlePercentage(b, user)
  })
}

func handlePercentage(b *tb.Bot, user *userInfo) {
    b.Handle(tb.OnText, func(m *tb.Message) {
      switch m.Text {
      case "1":
        user.Percentage = "25"
      case "2":
        user.Percentage = "35"
      case "3":
        user.Percentage = "50"
      case "4":
        user.Percentage = "70"
      default:
        b.Send(m.Sender, "Send number from 1 to 4")
        handlePercentage(b, user)
      }
      log.Println("Sender:", m.Sender, "Percentage:" , user.Percentage)
      handleStart(b, user)
    })
}
