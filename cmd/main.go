package main

import (
	barcode "MedunkaOpBarcode/pkg/Barcode"
	artist "MedunkaOpBarcode/pkg/CLIArtist"
	database "MedunkaOpBarcode/pkg/Database"
	env "MedunkaOpBarcode/pkg/Env" //create your own Env directory with env variables
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

	err := os.MkdirAll("/tmp/Products", os.ModePerm) //make dir to store product updates
	if err != nil {
		fmt.Println(err)
	}

	skladBois := telegrambot.User{Id: os.Getenv("botOwner")} //create user with id such that it is either group or user id generated by Telegram
	botHandler := telegrambot.Handler{Owner: skladBois}      //assign that user/group to the bot so that no one else talks to him
	botHandler.SetToken(os.Getenv("botToken"))

	var wg sync.WaitGroup //start the bot in a go routine so that he does not interfere with the rest of the program
	wg.Add(1)
	go func() {
		botHandler.StartBot()
	}()

	dbPort := typeconv.StringToInt(os.Getenv("dbPort"))
	dbConfig := database.DBConfig{
		Host:     os.Getenv("host"),
		Port:     dbPort,
		User:     os.Getenv("user"),
		Password: os.Getenv("password"),
		DBName:   os.Getenv("dbname"),
	}
	var postgresDBHandler database.PostgresDBHandler
	postgresDBHandler.GrabConfig(&dbConfig)
	postgresDBHandler.Connect()

	events.ReceiveFile(&botHandler, &postgresDBHandler, "/tmp/Products", "update.csv")

	serialPort := serialcommunication.OpenPort(&serial.Config{Name: "/dev/ttyAMA0", Baud: 9600})

	var barcodeReaderHandler barcode.ReaderHandler
	barcodeReaderHandler.GetPort(serialPort)

	for {

		barcodeAsByteArray := barcodeReaderHandler.ReadUntilDelimiter('\x0d')
		artist.ClearTerminal()

		barcodeAsString := typeconv.ByteArrayToString(barcodeAsByteArray[:len(barcodeAsByteArray)-1]) //cut out the delimiter and convert to string

		name, stock, price, unitOfMeasure, unitOfMeasureKoef := postgresDBHandler.QueryProductData(database.QueryProductDataSQL, barcodeAsString)

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
