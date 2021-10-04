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
		barcodeAsByteArray = barcodeAsByteArray[:len(barcodeAsByteArray)-1]
		barcodeAsInt, _ := strconv.ParseInt(string(barcodeAsByteArray), 10, 64)

		postgreHandler.Connect()
		price, mj, mjkoef := postgreHandler.QueryProductData(barcodeAsInt)
		postgreHandler.Disconnect()
		fmt.Println(price, mj, mjkoef)

	}
}