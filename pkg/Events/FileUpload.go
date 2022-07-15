package events

import (
	database "MedunkaOpBarcode/pkg/Database"
	essential "MedunkaOpBarcode/pkg/EssentialConfig"
	telegrambot "MedunkaOpBarcode/pkg/TelegramBot"
	"fmt"
	tb "gopkg.in/telebot.v3"
	"time"
)

func ReceiveFile(handler *telegrambot.Handler, db database.Database, conf *essential.Config) {
	handler.Bot.Handle(tb.OnDocument, func(c tb.Context) (err error) {
		if err != nil {
			return fmt.Errorf("unable to handle endpoint %s", err)
		}

		if !handler.OwnerVerify(c.Message().Chat.ID) {
			handler.SendText("unauthorised attempt to access bot: " + c.Message().Sender.Username)
			return fmt.Errorf("bot owner not verified")
		}

		if !handler.FileTypeVerify(c.Message().Document.FileName, ".csv") {
			handler.SendText("check file type")
			return fmt.Errorf("file type not valid")
		}
		fmt.Println("Downloading file...")
		handler.DownloadFile(&c.Message().Document.File, conf.PathToCSVUpdateFile, conf.CsvUpdateFileName)
		fmt.Println("Done!")
		time.Sleep(2 * time.Second)

		fmt.Println("Generating database table...")
		db.ExecuteStatement(database.GenerateDropExistingTableIfExistsSQL(conf.DatabaseTableName))
		db.ExecuteStatement(database.GenerateCreateTableSQL(conf.DatabaseTableName))
		db.ExecuteStatement(database.GenerateImportFromCSVToTableSQL(conf.DatabaseTableName, conf.PathToCSVUpdateFile, conf.CsvDelimiter))
		fmt.Println("Done!")
		return nil
	})
}
