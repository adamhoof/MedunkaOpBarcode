package events

import (
	database "MedunkaOPBarcode/pkg/Database"
	telegrambot "MedunkaOPBarcode/pkg/TelegramBot"
	"fmt"
	tb "gopkg.in/telebot.v3"
	"strings"
)

func FileUpload(handler *telegrambot.Handler, db database.Database) {
	handler.Bot.Handle(tb.OnDocument, func(c tb.Context) (err error) {
		if err != nil {
			fmt.Println(err)
		}
		if !strings.Contains(c.Message().Document.FileName, ".csv") {
			_, err = handler.Bot.Send(&handler.Owner, "Check file type")
			if err != nil {
				fmt.Println(err)
			}
		}
		return nil
	})
}
