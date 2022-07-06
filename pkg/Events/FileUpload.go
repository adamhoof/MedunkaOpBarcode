package events

import (
	"fmt"
	database "github.com/adamhoof/MedunkaOpBarcode/pkg/Database"
	telegrambot "github.com/adamhoof/MedunkaOpBarcode/pkg/TelegramBot"
	tb "gopkg.in/telebot.v3"
)

func FileUpload(handler *telegrambot.Handler, db database.Database) {
	handler.Bot.Handle(tb.OnDocument, func(c tb.Context) (err error) {
		if err != nil {
			return fmt.Errorf("unable to handle endpoint")
		}

		if !handler.OwnerVerify(c.Message().Chat.ID) {
			handler.SendText("unauthorised attempt to access bot: " + c.Message().Sender.Username)
			return fmt.Errorf("bot owner not verified")
		}

		if !handler.FileTypeVerify(c.Message().Document.FileName, ".csv") {
			handler.SendText("check file type")
			return fmt.Errorf("file type not valid")
		}
		handler.DownloadFile(&c.Message().Document.File, "Products/", "update.csv")

		db.ExecuteStatement(database.DropExistingTableSQL)
		db.ExecuteStatement(database.CreateTableSQL)
		db.ExecuteStatement(database.ImportFromCSVToTableSQL)
		return nil
	})
}
