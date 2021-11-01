package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gookit/color.v1"
	"strconv"
)

var actualPriceStyle = color.Style{color.FgRed, color.OpBold}
var defaultStyle = color.Style{color.FgLightWhite, color.OpItalic}

func main() {
	postgreHandler := PostgreSQLHandler{}
	serialHandler := SerialHandler{}
	barcodeController := BarcodeController{}
	formatter := Formatter{}

	postgreHandler.Connect()
	postgreHandler.DropTableIfExists()
	postgreHandler.CreateTable()
	postgreHandler.ImportFromCSV()

	serialHandler.PortConfig("/dev/ttyAMA0", 9600)
	serialHandler.OpenPort()

	barcodeController.CreateBarcodeReader(serialHandler.port)

	for {
		barcodeAsByteArray := barcodeController.Read()
		fmt.Print("\033[H\033[2J")

		barcodeAsByteArray = barcodeAsByteArray[:len(barcodeAsByteArray)-1]
		barcodeAsInt, _ := strconv.ParseInt(string(barcodeAsByteArray), 10, 64)

		name, stock, price, mj, mjkoef := postgreHandler.QueryProductData(barcodeAsInt)

		priceAsFloat := formatter.ToFloat(price)
		stringPricePerMj := formatter.ReturnAsString(priceAsFloat * mjkoef)

		formatter.PrintColoredText(defaultStyle, "\n"+"\n"+"\n"+"\n")
		formatter.PrintColoredText(defaultStyle, name)
		formatter.PrintColoredText(actualPriceStyle,
			"Cena za ks: "+
				price+"Kč"+
				"\n"+"\n"+"\n"+"\n"+"\n")

		formatter.PrintColoredText(defaultStyle, "Přepočet na měrnouj. ("+mj+"): "+
			stringPricePerMj+"Kč")
		formatter.PrintColoredText(defaultStyle, "\n")
		formatter.PrintColoredText(defaultStyle, "Stock: " + stock)
	}
}
