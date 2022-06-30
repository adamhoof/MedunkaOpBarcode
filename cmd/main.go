package main

import (
	"MedunkaOpBarcode/pkg/Barcode"
	"MedunkaOpBarcode/pkg/CLIArtist"
	"MedunkaOpBarcode/pkg/Database"
	"MedunkaOpBarcode/pkg/SerialCommunication"
	"MedunkaOpBarcode/pkg/TypeConvertor"
	_ "github.com/lib/pq"
	"gopkg.in/gookit/color.v1"
	"os"
	"strings"
)

const dropExistingTableSQL = `DROP TABLE IF EXISTS products;`
const createTableSQL = `CREATE TABLE products(Barcode text, name text, stock text, price text, mj text, mjkoef decimal);`
const importFromCSVToTableSQL = `COPY products FROM '/' DELIMITER ';' CSV HEADER;`
const queryProductDataSQL = `SELECT name, stock, price, mj, mjkoef FROM products WHERE Barcode = $1;`

var boldRed = color.Style{color.FgRed, color.OpBold}
var italicWhite = color.Style{color.FgLightWhite, color.OpItalic}

func main() {
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
	postgresDBHandler.ExecuteStatement(dropExistingTableSQL)
	postgresDBHandler.ExecuteStatement(createTableSQL)
	postgresDBHandler.ExecuteStatement(importFromCSVToTableSQL)

	serialPort := serialcommunication.OpenPort("/dev/ttyAMA0", 9600)

	var barcodeReaderHandler barcode.ReaderHandler
	barcodeReaderHandler.GetPort(serialPort)

	for {
		barcodeAsByteArray := barcodeReaderHandler.ReadUntilDelimiter('\x0d')
		artist.ClearTerminal()

		barcodeAsString := typeconv.ByteArrayToString(barcodeAsByteArray)

		name, stock, price, unitOfMeasure, unitOfMeasureKoef := postgresDBHandler.QueryProductData(queryProductDataSQL, barcodeAsString)

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
