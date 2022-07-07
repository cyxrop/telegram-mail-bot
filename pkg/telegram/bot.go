package telegram

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	botAPI *tgbotapi.BotAPI

	timeout int
}

func NewBot(token string, timeout int, debug bool) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("init bot: %w", err)
	}

	bot.Debug = debug

	return &Bot{
		botAPI:  bot,
		timeout: timeout,
	}, nil
}

func (b *Bot) SendMessage(chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := b.botAPI.Send(msg)

	return err
}

func (b *Bot) Updates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = b.timeout

	return b.botAPI.GetUpdatesChan(u)
}
