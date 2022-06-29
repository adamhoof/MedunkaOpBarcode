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
	"strconv"
	"strings"
)

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
	postgresDBHandler.ExecuteStatement(Database.DropExistingTableSQL)
	postgresDBHandler.ExecuteStatement(Database.CreateTableSQL)
	postgresDBHandler.ExecuteStatement(Database.ImportFromCSVToTableSQL)

	portConfig := SerialCommunication.CreatePortConfig("/dev/ttyAMA0", 9600)
	serialPort := SerialCommunication.OpenPort(portConfig)

	var reader *bufio.Reader
	Barcode.AssignPort(reader, serialPort)

	for {
		barcodeAsByteArray := Barcode.Read(reader)
		fmt.Print("\033[H\033[2J")

		barcodeAsByteArray = barcodeAsByteArray[:len(barcodeAsByteArray)-1]
		barcodeAsInt, _ := strconv.ParseInt(string(barcodeAsByteArray), 10, 64)
		barcodeAsString := Formatting.ToString(barcodeAsInt)

		name, stock, price, mj, mjkoef := postgresDBHandler.QueryProductData(barcodeAsString)

		formatedPrice := strings.ReplaceAll(price, ".00 Kč", "")
		stringPricePerMj := Formatting.ToString(Formatting.ToFloat(formatedPrice) * mjkoef)

		Formatting.PrintStyledText(Formatting.DefaultStyle, name)
		fmt.Println("")
		fmt.Println("")
		Formatting.PrintStyledText(Formatting.ActualPriceStyle,
			"Cena za ks: "+
				formatedPrice+"Kč"+
				"\n"+"\n")

		Formatting.PrintStyledText(Formatting.DefaultStyle, "Přepočet na měrnouj. ("+mj+"): "+
			stringPricePerMj+"Kč")
		Formatting.PrintStyledText(Formatting.DefaultStyle, "\n")
		Formatting.PrintStyledText(Formatting.DefaultStyle, "Stock: "+stock)
	}
}
