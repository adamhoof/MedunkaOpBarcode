package main

import (
	"MedunkaOpBarcode/pkg/Barcode"
	"MedunkaOpBarcode/pkg/Database"
	"MedunkaOpBarcode/pkg/Formatting"
	"MedunkaOpBarcode/pkg/SerialCommunication"
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"strings"
)

func main() {
	dbPort, err := strconv.Atoi(os.Getenv("dbPort"))
	if err != nil {
		panic(err)
	}
	dbConfig := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		os.Getenv("host"), dbPort, os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	var db *sql.DB
	Database.Connect(db, dbConfig)
	Database.ExecuteStatement(db, Database.DropExistingTableSQL)
	Database.ExecuteStatement(db, Database.CreateTableSQL)
	Database.ExecuteStatement(db, Database.ImportFromCSVToTableSQL)

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

		name, stock, price, mj, mjkoef := Database.QueryProductData(db, barcodeAsString)

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
