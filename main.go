package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gookit/color.v1"
	"strconv"
	"strings"
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
		barcodeAsString := formatter.ToString(barcodeAsInt)

		name, stock, price, mj, mjkoef := postgreHandler.QueryProductData(barcodeAsString)

		formatedPrice := strings.ReplaceAll(price, ".00 Kč", "")
		stringPricePerMj := formatter.ToString(formatter.ToFloat(formatedPrice) * mjkoef)

		formatter.PrintColoredText(defaultStyle, "\n"+"\n"+"\n"+"\n")
		formatter.PrintColoredText(defaultStyle, name+"\n"+"\n")
		formatter.PrintColoredText(actualPriceStyle,
			"Cena za ks: "+
				formatedPrice+"Kč"+
				"\n"+"\n"+"\n"+"\n"+"\n"+"\n")

		formatter.PrintColoredText(defaultStyle, "Přepočet na měrnouj. ("+mj+"): "+
			stringPricePerMj+"Kč")
		formatter.PrintColoredText(defaultStyle, "\n")
		formatter.PrintColoredText(defaultStyle, "Stock: "+stock)
	}
}
