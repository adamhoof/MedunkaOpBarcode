package telegrambot

import (
	database "MedunkaOpBarcode/pkg/Database"
	essential "MedunkaOpBarcode/pkg/EssentialConfig"
	"fmt"
	tb "gopkg.in/telebot.v3"
	"os"
	"os/exec"
	"time"
)

func (handler *Handler) OnReceiveFile(db database.Database, conf *essential.Config) (err error) {
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

		err = os.RemoveAll("/tmp/Products")
		if err != nil {
			return err
		}
		err = os.MkdirAll("/tmp/Products", os.ModePerm)
		if err != nil {
			return err
		}

		fmt.Println("Downloading file...")
		err = handler.DownloadFile(&c.Message().Document.File, conf.PathToCSVUpdateFile, c.Message().Document.FileName)
		if err != nil {
			return err
		}
		fmt.Println("Done!")
		time.Sleep(2 * time.Second)

		switch fileType {
		case ".xlsx":
			exec.Command("xlsx2csv", "/tmp/Products/"+c.Message().Document.FileName+".xlsx", "-d", ";", "/tmp/Products/"+conf.CsvUpdateFileName)
		case ".xls":
			exec.Command("xlsx2csv", "/tmp/Products/"+c.Message().Document.FileName+".xls", "-d", ";", "/tmp/Products/"+conf.CsvUpdateFileName)
		default:
			fmt.Println("This should never execute, or file type verification is broken")
		}
		fmt.Println("File type: ", fileType)

		fmt.Println("Generating database table...")
		err = db.ExecuteStatement(database.GenerateDropExistingTableIfExistsSQL(conf.DatabaseTableName))
		if err != nil {
			return err
		}
		err = db.ExecuteStatement(database.GenerateCreateTableSQL(conf.DatabaseTableName))
		if err != nil {
			return err
		}
		err = db.ExecuteStatement(database.GenerateImportFromCSVToTableSQL(conf.DatabaseTableName, conf.PathToCSVUpdateFile, conf.CsvUpdateFileName, conf.CsvDelimiter))
		if err != nil {
			return err
		}
		fmt.Println("Done!")
		return err
	})
	return err
}
