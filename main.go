package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
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
		rawBarcode, err := barcodeController.barcode.ReadBytes('\x0d')
		if err != nil {
			panic(err)
		}

		stringBarcodeOutput := string(rawBarcode)
		stringBarcodeOutput = strings.Replace(stringBarcodeOutput, "\r", "", -1)

		var formatedBarcode int64
		if i, convErr := strconv.ParseInt(stringBarcodeOutput, 10, 64); convErr == nil {
			formatedBarcode = i
		}
		fmt.Println(formatedBarcode)
	}
}