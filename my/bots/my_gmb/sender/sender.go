package sender

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

type Sender struct {
	token  string
	chatId int
	bot    *tgbotapi.BotAPI
}

func TgSender(
	token string,
	chatId int,
) (*Sender, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("error creating new tg bot: %w", err)
	}

	return &Sender{
		bot:    bot,
		token:  token,
		chatId: chatId,
	}, nil
}

func (s *Sender) Send(message string) error {
	//bot, err := tgbotapi.NewBotAPI(s.token)
	//if err != nil {
	//	return fmt.Errorf("error creating new tg bot: %w", err)
	//}

	msg := tgbotapi.NewMessage(int64(s.chatId), message)
	msg.ParseMode = "Markdown"

	_, err := s.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("message sending error: %w", err)
	}

	return nil
}

func (s *Sender) SendVideo() error {
	photo := tgbotapi.NewVideo(int64(s.chatId), tgbotapi.FileURL(fmt.Sprintf("https://cataas.com/cat/gif?%v", time.Now())))
	if _, err := s.bot.Send(photo); err != nil {
		return fmt.Errorf("message sending video error: %w", err)
	}

	return nil
}
