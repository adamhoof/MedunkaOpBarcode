package telegrambot

import (
	"fmt"
	tb "gopkg.in/telebot.v3"
	"strings"
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

func (handler *Handler) OwnerVerify(id int64) bool {
	if id != handler.Owner.IDAsInt() {
		return false
	}
	return true
}

func (handler *Handler) SendText(text string) {
	_, err := handler.Bot.Send(&handler.Owner, text)
	if err != nil {
		fmt.Println(err)
	}
}

func (handler *Handler) FileTypeVerify(fileName string, fileType string) bool {
	if !strings.Contains(fileName, fileType) {
		return false
	}
	return true
}

func (handler *Handler) DownloadFile(file *tb.File, path string, chosenName string) {
	err := handler.Bot.Download(file, path+chosenName)
	if err != nil {
		fmt.Println(err)
	}
}
