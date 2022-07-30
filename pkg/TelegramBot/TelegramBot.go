package telegrambot

import (
	typeconv "MedunkaOpBarcode/pkg/TypeConversion"
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
	if id != typeconv.StringToInt64(handler.Owner.Id) {
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

func (handler *Handler) FileTypeVerify(fileName string, validFileTypes []string) (string, bool) {
	for _, value := range validFileTypes {
		if strings.Contains(fileName, value) {
			return value, true
		}
		continue
	}
	return "", false
}

func (handler *Handler) DownloadFile(file *tb.File, path string, name string) (err error) {
	err = handler.Bot.Download(file, path+name)
	if err != nil {
		return fmt.Errorf("failed to download file: %s", err)
	}
	return nil
}
