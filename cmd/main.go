package main

import (
	barcode "github.com/adamhoof/MedunkaOpBarcode/pkg/Barcode"
	artist "github.com/adamhoof/MedunkaOpBarcode/pkg/CLIArtist"
	database "github.com/adamhoof/MedunkaOpBarcode/pkg/Database"
	env "github.com/adamhoof/MedunkaOpBarcode/pkg/Env"
	events "github.com/adamhoof/MedunkaOpBarcode/pkg/Events"
	serialcommunication "github.com/adamhoof/MedunkaOpBarcode/pkg/SerialCommunication"
	telegrambot "github.com/adamhoof/MedunkaOpBarcode/pkg/TelegramBot"
	typeconv "github.com/adamhoof/MedunkaOpBarcode/pkg/TypeConversion"
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

	skladBois := telegrambot.User{Id: "-1001671432440"}
	botHandler := telegrambot.Handler{Owner: skladBois}
	botHandler.SetToken(os.Getenv("botToken"))

	events.FileUpload(&botHandler, &database.PostgresDBHandler{})

	var wg sync.WaitGroup
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
	postgresDBHandler.ExecuteStatement(database.DropExistingTableSQL)
	postgresDBHandler.ExecuteStatement(database.CreateTableSQL)
	postgresDBHandler.ExecuteStatement(database.ImportFromCSVToTableSQL)

	serialPort := serialcommunication.OpenPort(&serial.Config{Name: "/dev/ttyAMA0", Baud: 9600})

	var barcodeReaderHandler barcode.ReaderHandler
	barcodeReaderHandler.GetPort(serialPort)

	for {
		barcodeAsByteArray := barcodeReaderHandler.ReadUntilDelimiter('\x0d')
		artist.ClearTerminal()

		barcodeAsString := typeconv.ByteArrayToString(barcodeAsByteArray)

		name, stock, price, unitOfMeasure, unitOfMeasureKoef := postgresDBHandler.QueryProductData(database.QueryProductDataSQL, barcodeAsString)

		strPriceWithoutSuffix := strings.ReplaceAll(price, ".00 Kč", "")
		strPricePerMj := typeconv.FloatToString(typeconv.StringToFloat(strPriceWithoutSuffix) * unitOfMeasureKoef)

		artist.PrintStyledText(italicWhite, name)
		artist.PrintSpaces(2)
		artist.PrintStyledText(boldRed,
			"Cena za ks: "+
				strPriceWithoutSuffix+"Kč"+
				"\n"+"\n") //TODO replace?

		artist.PrintStyledText(italicWhite, "Přepočet na měrnouj. ("+unitOfMeasure+"): "+
			strPricePerMj+"Kč")
		artist.PrintStyledText(italicWhite, "\n") //TODO replace?
		artist.PrintStyledText(italicWhite, "Stock: "+stock)
	}
}
