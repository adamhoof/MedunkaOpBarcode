package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

func main() {
	postgreHandler := PostgreSQLHandler{}
	postgreHandler.Connect()
	postgreHandler.DropTableIfExists()
	postgreHandler.CreateTable()
	postgreHandler.ImportFromCSV()
	postgreHandler.Disconnect()

	barcodeController := BarcodeController{}
	serialHandler := SerialHandler{}
	serialHandler.PortConfig("/dev/ttyAMA0", 9600)
	serialHandler.OpenPort()
	barcodeController.CreateBarcodeReader(serialHandler.port)

	for {
		barcodeOutput, err := barcodeController.barcode.ReadBytes('\x0d')
		if err != nil {
			panic(err)
		}
		fmt.Println(barcodeOutput)
		stringBarcodeOutput := string(barcodeOutput)
		stringBarcodeOutput = strings.Replace(stringBarcodeOutput, "\r", "", -1)
	}
}