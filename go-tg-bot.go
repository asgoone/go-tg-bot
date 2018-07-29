package main

import (
"log"
"github.com/Syfaro/telegram-bot-api"
"fmt"
"os"
"encoding/json"
)

type Config struct {
	TelegramBotToken string
}

func main() {

	file, _ := os.Open("config.json")
		decoder := json.NewDecoder(file)
		configuration := Config{}
		err := decoder.Decode(&configuration)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(configuration.TelegramBotToken)

    // указываем токен для доступу к боту берем из конфига 
		bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
		if err != nil {
			log.Panic(err)
		}

		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil {
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}