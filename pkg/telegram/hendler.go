package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Infof("Message form %v", message.From.UserName)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")

	switch message.Command() {

	case "status":
		msg, err := b.status(message)
		if err != nil {
			log.Error(err)
			return err
		}
		if _, err := b.bot.Send(msg); err != nil {
			log.Error(err)
			return err
		}
	case "user_of_group":
		msg, err := b.members(message)
		if err != nil {
			log.Error(err)
			return err
		}
		if _, err := b.bot.Send(msg); err != nil {
			log.Error(err)
			return err
		}
	default:
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

	}
	return nil
}

func (b *Bot) status(message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	valutes, err := b.newRequest()
	if err != nil {
		log.Error(err)
		return tgbotapi.MessageConfig{}, err
	}
	answer := b.getAnswer(valutes)
	msg := tgbotapi.NewMessage(message.Chat.ID, answer)
	return msg, nil
}

func (b *Bot) members(message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Chat.FirstName+" "+message.Chat.LastName)
	return msg, nil
}
