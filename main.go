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
		if formatedBarcode, err := strconv.Atoi(stringBarcodeOutput); err == nil {
			fmt.Printf("i=%d, type: %T\n", formatedBarcode, formatedBarcode)
		}
	}
}