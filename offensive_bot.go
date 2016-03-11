package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
    "os"
    "fmt"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Args[1])
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        replytext := fmt.Sprintf("Oi, %s", update.Message.From.FirstName)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, replytext)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
