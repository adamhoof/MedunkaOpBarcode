package telegrambot

import (
	"fmt"
	tb "gopkg.in/telebot.v3"
	"time"
)

type Handler struct {
	Bot   *tb.Bot
	Owner User
}

func (handler *Handler) SetToken(token string) {
	var err error
	handler.Bot, err = tb.NewBot(tb.Settings{
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

func (handler *Handler) StartBot() {
	handler.Bot.Start()
}
