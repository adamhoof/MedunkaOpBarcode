package main

import (
	"MedunkaOpBarcode/pkg/Barcode"
	"MedunkaOpBarcode/pkg/Database"
	"MedunkaOpBarcode/pkg/Formatting"
	"MedunkaOpBarcode/pkg/SerialCommunication"
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

const dropExistingTableSQL = `DROP TABLE IF EXISTS products;`
const createTableSQL = `CREATE TABLE products(barcode text, name text, stock text, price text, mj text, mjkoef decimal);`
const importFromCSVToTableSQL = `COPY products FROM '/' DELIMITER ';' CSV HEADER;`
const queryProductDataSQL = `SELECT name, stock, price, mj, mjkoef FROM products WHERE barcode = $1;`

func main() {
	dbPort := Formatting.StringToInt(os.Getenv("dbPort"))
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

		barcodeAsString := Formatting.ByteArrayToString(barcodeAsByteArray)

		name, stock, price, mj, mjKoef := postgresDBHandler.QueryProductData(queryProductDataSQL, barcodeAsString)

		formattedPrice := strings.ReplaceAll(price, ".00 Kč", "")
		stringPricePerMj := Formatting.FloatToString(Formatting.StringToFloat(formattedPrice) * mjKoef)

		Formatting.PrintStyledText(Formatting.Default, name)
		Formatting.PrintSpaces(2)
		Formatting.PrintStyledText(Formatting.BoldRed,
			"Cena za ks: "+
				formattedPrice+"Kč"+
				"\n"+"\n")

		Formatting.PrintStyledText(Formatting.Default, "Přepočet na měrnouj. ("+mj+"): "+
			stringPricePerMj+"Kč")
		Formatting.PrintStyledText(Formatting.Default, "\n")
		Formatting.PrintStyledText(Formatting.Default, "Stock: "+stock)
	}
}
