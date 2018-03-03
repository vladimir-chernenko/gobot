package main

import (
	"log"
  "regexp"
	tb "gopkg.in/tucnak/telebot.v2"
)

type userInfo struct {
  phone string
  receiverName string
  senderName string
  currency string
  percentage string
}

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
      user.phone = m.Text
      log.Println("Sender:", m.Sender, "Phone:" , user.phone)
      b.Send(m.Sender, "Ok, now I need receiver's name")
      handleReceiverName(b, user)
    } else {
      b.Send(m.Sender, "Number is not valid. Try again.")
      handlePhoneNumber(b, user)
    }
  })
}

func handleReceiverName(b *tb.Bot, user *userInfo)  {
  b.Handle(tb.OnText, func(m *tb.Message) {
    if m.Text != "" {
      user.receiverName = m.Text
      log.Println("Sender:", m.Sender, "Receiver name:" , user.receiverName)
      b.Send(m.Sender, "Send me now your name")
      handleSenderName(b, user)
    } else {
      handleReceiverName(b, user)
    }
  })
}

func handleSenderName(b *tb.Bot, user *userInfo)  {
  b.Handle(tb.OnText, func(m *tb.Message) {
    if m.Text != "" {
      user.senderName = m.Text
      log.Println("Sender:", m.Sender, "Sender name:" , user.senderName)
      b.Send(m.Sender, "Now choose currency. Send me number that stands for currency you want")
      b.Send(m.Sender, "1 Bitcoin\n2 Ethereum\n3 Litecoin\n4 Bitcoin Cash")
      handleCurrency(b, user)
    } else {
      handleReceiverName(b, user)
    }
  })
}

func handleCurrency(b *tb.Bot, user *userInfo) {
  b.Handle(tb.OnText, func(m *tb.Message) {
    switch m.Text {
    case "1":
      user.currency = "Bitcoin"
    case "2":
      user.currency = "Ethereum"
    case "3":
      user.currency = "Litecoin"
    case "4":
      user.currency = "Bitcoin Cash"
    default:
      b.Send(m.Sender, "Send number from 1 to 4")
      handleCurrency(b, user)
    }
    log.Println("Sender:", m.Sender, "Currency:" , user.currency)
    b.Send(m.Sender, "And the last thing. Send message when it drops by")
    b.Send(m.Sender, "1 25%\n2 35%\n3 50%\n4 70%")
    handlePercentage(b, user)
  })
}

func handlePercentage(b *tb.Bot, user *userInfo) {
    b.Handle(tb.OnText, func(m *tb.Message) {
      switch m.Text {
      case "1":
        user.percentage = "25"
      case "2":
        user.percentage = "35"
      case "3":
        user.percentage = "50"
      case "4":
        user.percentage = "70"
      default:
        b.Send(m.Sender, "Send number from 1 to 4")
        handlePercentage(b, user)
      }
      log.Println("Sender:", m.Sender, "Percentage:" , user.percentage)
      handleStart(b, user)
    })
}
