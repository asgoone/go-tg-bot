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

// получаем из config.json токен
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
		// Записываем токен в переменную configuration.TelegramBotToken
		fmt.Println(configuration.TelegramBotToken) 

    	// указываем токен для доступу к боту берем из конфига 
		bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
		if err != nil {
			log.Panic(err)
		}

		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)

		// u - структура с конфигом для получения апдейтов
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		// используя конфиг u создаем канал в который будут прилетать новые сообщения
		updates, err := bot.GetUpdatesChan(u)


		// в канал updates прилетают структуры типа Update
		// вычитываем их и обрабатываем
		for update := range updates {
			// reply := "Не знаю что сказать"
			if update.Message == nil {
				continue
			}

			// логируем от кого какое сообщение пришло
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// создаем ответное сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			// ??
			msg.ReplyToMessageID = update.Message.MessageID
			
			// и отправляем его
			bot.Send(msg)
		}
	}