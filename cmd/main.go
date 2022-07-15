package main

import (
	barcode "MedunkaOpBarcode/pkg/Barcode"
	artist "MedunkaOpBarcode/pkg/CLIArtist"
	database "MedunkaOpBarcode/pkg/Database"
	env "MedunkaOpBarcode/pkg/Env" //you do not have to use your own Env file if you don't publish it anywhere, enter the credentials into the EssentialConfig
	essential "MedunkaOpBarcode/pkg/EssentialConfig"
	events "MedunkaOpBarcode/pkg/Events"
	serialcommunication "MedunkaOpBarcode/pkg/SerialCommunication"
	telegrambot "MedunkaOpBarcode/pkg/TelegramBot"
	typeconv "MedunkaOpBarcode/pkg/TypeConversion"
	"fmt"
	"github.com/tarm/serial"
	"gopkg.in/gookit/color.v1"
	"os"
	"strings"
	"sync"
)

var boldRed = color.Style{color.FgRed, color.OpBold}
var italicWhite = color.Style{color.FgLightWhite, color.OpItalic}

func main() {
	env.SetEnv()

	conf := essential.Config{ //configure your own values based on your hardware, privileges, settings etc., those values are going to be used throughout the program
		PathToCSVUpdateFile:                "/tmp/Products",
		CsvUpdateFileName:                  "update.csv",
		CsvDelimiter:                       ";",
		DatabaseTableName:                  "products",
		SerialPortName:                     "/dev/ttyAMA0",
		SerialPortBaudRate:                 9600,
		BarcodeReadingTerminationDelimiter: '\x0d',
		DbHost:                             os.Getenv("host"),
		DbName:                             os.Getenv("dbname"),
		DbPort:                             typeconv.StringToInt(os.Getenv("dbPort")),
		DbUser:                             os.Getenv("user"),
		DbUserPassword:                     os.Getenv("password"),
		TelegramBotToken:                   os.Getenv("botToken"),
		TelegramBotOwner:                   os.Getenv("botOwner"),
	}

	err := os.MkdirAll(conf.PathToCSVUpdateFile, os.ModePerm) //make dir to store product updates
	if err != nil {
		fmt.Println(err)
	}

	skladBois := telegrambot.User{Id: conf.TelegramBotOwner} //create user with id such that it is either group or user id generated by Telegram
	botHandler := telegrambot.Handler{Owner: skladBois}      //assign that user/group to the bot so that no one else talks to him
	botHandler.SetToken(conf.TelegramBotToken)

	var wg sync.WaitGroup //start the bot in a go routine so that he does not interfere with the rest of the program
	wg.Add(1)
	go func() {
		botHandler.StartBot()
	}()

	dbConnectionConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		conf.DbHost,
		conf.DbPort,
		conf.DbUser,
		conf.DbUserPassword,
		conf.DbName)

	var postgresDBHandler database.PostgresDBHandler
	postgresDBHandler.Connect(&dbConnectionConfig)

	events.ReceiveFile(&botHandler, &postgresDBHandler, &conf)

	serialPort := serialcommunication.OpenPort(&serial.Config{Name: conf.SerialPortName, Baud: conf.SerialPortBaudRate})

	var barcodeReaderHandler barcode.ReaderHandler
	barcodeReaderHandler.GetPort(serialPort)

	for {
		barcodeAsByteArray := barcodeReaderHandler.ReadUntilDelimiter(conf.BarcodeReadingTerminationDelimiter)
		artist.ClearTerminal()

		barcodeAsString := typeconv.ByteArrayToString(barcodeAsByteArray[:len(barcodeAsByteArray)-1]) //cut out the delimiter and convert to string

		name, stock, price, unitOfMeasure, unitOfMeasureKoef := postgresDBHandler.QueryProductData(database.GenerateQueryProductDataSQL(conf.DatabaseTableName), barcodeAsString)

		strPriceWithoutSuffix := strings.ReplaceAll(price, ".00 Kč", "")
		strPricePerMj := typeconv.FloatToString(typeconv.StringToFloat(strPriceWithoutSuffix) * unitOfMeasureKoef)

		artist.PrintStyledText(italicWhite, name)
		artist.PrintSpaces(2)
		artist.PrintStyledText(boldRed, fmt.Sprintf("Cena za ks: %s Kč", strPriceWithoutSuffix))
		artist.PrintSpaces(2)

		if unitOfMeasure == "" {
			artist.PrintStyledText(italicWhite, "Stock: "+stock)
			continue
		}
		artist.PrintStyledText(italicWhite, fmt.Sprintf("Přepočet na %s: %s Kč", unitOfMeasure, strPricePerMj))
		artist.PrintSpaces(1)
		artist.PrintStyledText(italicWhite, "Stock: "+stock)
	}
}
