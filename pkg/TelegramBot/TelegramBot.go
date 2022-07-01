package telegrambot

import (
	"fmt"
	tb "gopkg.in/telebot.v3"
	"time"
)

type Handler struct {
	bot *tb.Bot
}

func (handler *Handler) SetToken(token string) {
	var err error
	handler.bot, err = tb.NewBot(tb.Settings{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("handler token valid")
}

func (handler *Handler) OnUploadFileAction() {
	handler.bot.Handle(tb.OnDocument, func(c tb.Context) error {
		return nil //TODO implement
	})
}

func (handler *Handler) StartBot() {
	handler.bot.Start()
}
