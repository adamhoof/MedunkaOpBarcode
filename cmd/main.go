package main

import (
	"MedunkaOpBarcode/pkg/Barcode"
	"MedunkaOpBarcode/pkg/Database"
	"MedunkaOpBarcode/pkg/SerialCommunication"
	"MedunkaOpBarcode/pkg/TextFormatting"
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gookit/color.v1"
	"strconv"
	"strings"
)

var actualPriceStyle = color.Style{color.FgRed, color.OpBold}
var defaultStyle = color.Style{color.FgLightWhite, color.OpItalic}

func main() {
	postgreHandler := Database.PostgreSQLHandler{}
	formatter := TextFormatting.Formatter{}

	postgreHandler.Connect()
	postgreHandler.DropTableIfExists()
	postgreHandler.CreateTable()
	postgreHandler.ImportFromCSV()

	portConfig := SerialCommunication.CreatePortConfig("/dev/ttyAMA0", 9600)
	port := SerialCommunication.OpenPort(portConfig)

	var reader *bufio.Reader
	Barcode.AssignPort(reader, port)

	for {
		barcodeAsByteArray := Barcode.Read(reader)
		fmt.Print("\033[H\033[2J")

		barcodeAsByteArray = barcodeAsByteArray[:len(barcodeAsByteArray)-1]
		barcodeAsInt, _ := strconv.ParseInt(string(barcodeAsByteArray), 10, 64)
		barcodeAsString := formatter.ToString(barcodeAsInt)

		name, stock, price, mj, mjkoef := postgreHandler.QueryProductData(barcodeAsString)

		formatedPrice := strings.ReplaceAll(price, ".00 Kč", "")
		stringPricePerMj := formatter.ToString(formatter.ToFloat(formatedPrice) * mjkoef)

		formatter.PrintColoredText(defaultStyle, name)
		fmt.Println("")
		fmt.Println("")
		formatter.PrintColoredText(actualPriceStyle,
			"Cena za ks: "+
				formatedPrice+"Kč"+
				"\n"+"\n")

		formatter.PrintColoredText(defaultStyle, "Přepočet na měrnouj. ("+mj+"): "+
			stringPricePerMj+"Kč")
		formatter.PrintColoredText(defaultStyle, "\n")
		formatter.PrintColoredText(defaultStyle, "Stock: "+stock)
	}
}
