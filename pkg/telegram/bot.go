package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	updates := b.initUpdatesChan()
	b.handleUpdates(updates)
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}
		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) newRequest() (map[string]interface{}, error) {
	data := make(map[string]interface{})
	client := http.Client{}
	reuest, err := http.NewRequest(http.MethodGet, "https://www.cbr-xml-daily.ru/daily_json.js", nil)
	if err != nil {
		log.Fatalf("Get Daily err: %v", err)
		return nil, err
	}
	res, err := client.Do(reuest)
	if err != nil {
		log.Printf("Error", err)
		return nil, err
	} else {
		defer func() {
			if err := res.Body.Close(); err != nil {
				log.Printf("close response body error: %v", err)
			}
		}()
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
	}
	_ = json.Unmarshal(body, &data)
	valute, ok := data["Valute"]
	if !ok {
		log.Error("Ne ok")
	}
	valutes, _ := valute.(map[string]interface{})
	return valutes, nil
}

func (b *Bot) getAnswer(valute map[string]interface{}) string {
	keys := []string{"USD", "EUR", "IDR", "THB"}
	var answer string
	for _, key := range keys {
		result, _ := valute[key]
		resultMap, _ := result.(map[string]interface{})
		resIDR, _ := resultMap["Value"]
		nameIDR, _ := resultMap["Name"]
		answer = answer + fmt.Sprint(nameIDR) + "\n" + fmt.Sprint(resIDR) + "\n"
	}

	return answer
}
