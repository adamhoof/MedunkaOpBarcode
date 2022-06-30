package main

import (
	"MedunkaOpBarcode/pkg/Barcode"
	"MedunkaOpBarcode/pkg/CLIArtist"
	"MedunkaOpBarcode/pkg/Database"
	"MedunkaOpBarcode/pkg/SerialCommunication"
	"MedunkaOpBarcode/pkg/TypeConvertor"
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gookit/color.v1"
	"os"
	"strings"
)

const dropExistingTableSQL = `DROP TABLE IF EXISTS products;`
const createTableSQL = `CREATE TABLE products(barcode text, name text, stock text, price text, mj text, mjkoef decimal);`
const importFromCSVToTableSQL = `COPY products FROM '/' DELIMITER ';' CSV HEADER;`
const queryProductDataSQL = `SELECT name, stock, price, mj, mjkoef FROM products WHERE barcode = $1;`

var boldRed = color.Style{color.FgRed, color.OpBold}
var italicWhite = color.Style{color.FgLightWhite, color.OpItalic}

func main() {
	dbPort := TypeConvertor.StringToInt(os.Getenv("dbPort"))
	dbConfig := Database.DBConfig{
		Host:     os.Getenv("host"),
		Port:     dbPort,
		User:     os.Getenv("user"),
		Password: os.Getenv("password"),
		DBName:   os.Getenv("dbname"),
	}
	var postgresDBHandler Database.PostgresDBHandler
	postgresDBHandler.GrabConfig(&dbConfig)
	postgresDBHandler.Connect()
	postgresDBHandler.ExecuteStatement(dropExistingTableSQL)
	postgresDBHandler.ExecuteStatement(createTableSQL)
	postgresDBHandler.ExecuteStatement(importFromCSVToTableSQL)

	portConfig := SerialCommunication.CreatePortConfig("/dev/ttyAMA0", 9600)
	serialPort := SerialCommunication.OpenPort(portConfig)

	var reader *bufio.Reader
	Barcode.AssignPort(reader, serialPort)

	for {
		barcodeAsByteArray := Barcode.Read(reader)
		fmt.Print("\033[H\033[2J")

		barcodeAsString := TypeConvertor.ByteArrayToString(barcodeAsByteArray)

		name, stock, price, mj, mjKoef := postgresDBHandler.QueryProductData(queryProductDataSQL, barcodeAsString)

		strPriceWithoutSuffix := strings.ReplaceAll(price, ".00 Kč", "")
		strPricePerMj := TypeConvertor.FloatToString(TypeConvertor.StringToFloat(strPriceWithoutSuffix) * mjKoef)

		CLIArtist.PrintStyledText(italicWhite, name)
		CLIArtist.PrintSpaces(2)
		CLIArtist.PrintStyledText(boldRed,
			"Cena za ks: "+
				strPriceWithoutSuffix+"Kč"+
				"\n"+"\n")

		CLIArtist.PrintStyledText(italicWhite, "Přepočet na měrnouj. ("+mj+"): "+
			strPricePerMj+"Kč")
		CLIArtist.PrintStyledText(italicWhite, "\n")
		CLIArtist.PrintStyledText(italicWhite, "Stock: "+stock)
	}
}
