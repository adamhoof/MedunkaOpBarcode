package main

import (
	database "MedunkaOpBarcode/pkg/Database"
	env "MedunkaOpBarcode/pkg/Env" //create your own Env directory with env variables
	events "MedunkaOpBarcode/pkg/Events"
	telegrambot "MedunkaOpBarcode/pkg/TelegramBot"
	typeconv "MedunkaOpBarcode/pkg/TypeConversion"
	"fmt"
	"gopkg.in/gookit/color.v1"
	"os"
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

	skladBois := telegrambot.User{Id: os.Getenv("botOwner")}
	botHandler := telegrambot.Handler{Owner: skladBois} //set owner for bot
	botHandler.SetToken(os.Getenv("botToken"))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		botHandler.StartBot() //start the bot in a go routine so that he does not interfere with the rest of the program
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

	events.ReceiveFile(&botHandler, &postgresDBHandler)

	/*serialPort := serialcommunication.OpenPort(&serial.Config{Name: "/dev/ttyAMA0", Baud: 9600})

	var barcodeReaderHandler barcode.ReaderHandler
	barcodeReaderHandler.GetPort(serialPort)

	for {

		barcodeAsByteArray := barcodeReaderHandler.ReadUntilDelimiter('\x0d')
		artist.ClearTerminal()

		//cut out the delimiter and convert to string
		barcodeAsString := typeconv.ByteArrayToString(barcodeAsByteArray[:len(barcodeAsByteArray)-1])

		name, stock, price, unitOfMeasure, unitOfMeasureKoef := postgresDBHandler.QueryProductData(database.QueryProductDataSQL, barcodeAsString)

		strPriceWithoutSuffix := strings.ReplaceAll(price, ".00 Kč", "")
		strPricePerMj := typeconv.FloatToString(typeconv.StringToFloat(strPriceWithoutSuffix) * unitOfMeasureKoef)

		artist.PrintStyledText(italicWhite, name)
		artist.PrintSpaces(2)
		artist.PrintStyledText(boldRed,
			"Cena za ks: "+
				strPriceWithoutSuffix+"Kč"+
				"\n"+"\n") //TODO replace?

		if unitOfMeasure == "" {
			artist.PrintStyledText(italicWhite, "Stock: "+stock)
			continue
		}
		artist.PrintStyledText(italicWhite, "Přepočet na ("+unitOfMeasure+"): "+
			strPricePerMj+"Kč")
		artist.PrintStyledText(italicWhite, "\n") //TODO replace?
		artist.PrintStyledText(italicWhite, "Stock: "+stock)
	}*/
}
