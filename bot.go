package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type newsBot struct {
	*tgbotapi.BotAPI
	channelID int64
}

func newNewsBot(config *Config) *newsBot {
	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &newsBot{
		BotAPI:    bot,
		channelID: config.ChannelID,
	}
}

func (newsbot *newsBot) sendNews(newsList []article) {
	if newsList == nil || len(newsList) == 0 {
		return
	}

	for _, news := range newsList {
		newsbot.Send(tgbotapi.NewMessage(newsbot.channelID, news.String()))
		time.Sleep(3 * time.Second)
	}
}
