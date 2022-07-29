package events

import (
	database "MedunkaOpBarcode/pkg/Database"
	essential "MedunkaOpBarcode/pkg/EssentialConfig"
	telegrambot "MedunkaOpBarcode/pkg/TelegramBot"
	"fmt"
	tb "gopkg.in/telebot.v3"
	"os/exec"
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

		fileType, isValid := handler.FileTypeVerify(c.Message().Document.FileName, []string{".csv", ".xlsx", ".xls"})
		if !isValid {
			handler.SendText("check file type")
			return fmt.Errorf("file type not valid")
		}

		fmt.Println("Downloading file...")
		handler.DownloadFile(&c.Message().Document.File, conf.PathToCSVUpdateFile, c.Message().Document.FileName)
		fmt.Println("Done!")
		time.Sleep(100 * time.Millisecond)

		switch fileType {
		case ".xlsx":
			exec.Command("xlsx2csv", "/tmp/Products/"+c.Message().Document.FileName+".xlsx", "-d", ";", "/tmp/Products/"+conf.CsvUpdateFileName)
		case ".xls":
			exec.Command("xlsx2csv", "/tmp/Products/"+c.Message().Document.FileName+".xls", "-d", ";", "/tmp/Products/"+conf.CsvUpdateFileName)
		default:
			fmt.Println("This should never execute, or file type verification is incorrect")
		}
		fmt.Println("File type: ", fileType)

		fmt.Println("Generating database table...")
		db.ExecuteStatement(database.GenerateDropExistingTableIfExistsSQL(conf.DatabaseTableName))
		db.ExecuteStatement(database.GenerateCreateTableSQL(conf.DatabaseTableName))
		db.ExecuteStatement(database.GenerateImportFromCSVToTableSQL(conf.DatabaseTableName, conf.PathToCSVUpdateFile, conf.CsvUpdateFileName, conf.CsvDelimiter))
		fmt.Println("Done!")
		return nil
	})
}
