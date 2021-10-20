package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gookit/color.v1"
	"strconv"
)

var actualPriceColor = color.Style{color.FgRed, color.OpBold, color.BgLightWhite}
var unitOfMessureColor = color.Style{color.FgLightWhite}

func main() {
	postgreHandler := PostgreSQLHandler{}
	serialHandler := SerialHandler{}
	barcodeController := BarcodeController{}
	formatter := Formatter{}

	postgreHandler.Connect()
	postgreHandler.DropTableIfExists()
	postgreHandler.CreateTable()
	postgreHandler.ImportFromCSV()
	postgreHandler.Disconnect()

	serialHandler.PortConfig("/dev/ttyAMA0", 9600)
	serialHandler.OpenPort()

	barcodeController.CreateBarcodeReader(serialHandler.port)

	for {
		barcodeAsByteArray := barcodeController.Read()
		fmt.Print("\033[H\033[2J")

		barcodeAsByteArray = barcodeAsByteArray[:len(barcodeAsByteArray)-1]
		barcodeAsInt, _ := strconv.ParseInt(string(barcodeAsByteArray), 10, 64)

		postgreHandler.Connect()
		price, mj, mjkoef := postgreHandler.QueryProductData(barcodeAsInt)
		postgreHandler.Disconnect()

		stringPrice := formatter.ReturnAsString(price)
		stringPricePerMj := formatter.ReturnAsString(float32(price) * mjkoef)

		formatter.PrintColoredText(actualPriceColor, "\n"+"\n"+"\n"+"\n"+"\n"+
			"Cena za kus: "+
			stringPrice+" Kč"+
			"\n"+"\n"+"\n")

		formatter.PrintColoredText(unitOfMessureColor, "Přepočet na měrnou j. ("+mj+"): "+
			stringPricePerMj+" Kč")

	}
}
