package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

func main() {
	postgreHandler := PostgreSQLHandler{}
	serialHandler := SerialHandler{}
	barcodeController := BarcodeController{}

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

		pricePerMj := int32(float32(price)*mjkoef)

		fmt.Println("\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "Cena za kus: ")
		fmt.Print(price)
		fmt.Print(" Kč")
		fmt.Println("\n" +"\n" +"\n")
		fmt.Print("Přepočet na měrnou j. (" + mj + "): ")
		fmt.Print(pricePerMj)
		fmt.Print(" Kč")
	}
}
