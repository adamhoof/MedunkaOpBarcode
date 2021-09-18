package main

import (
	"fmt"
	"strings"
)

func main() {
	apiHandler := APIHandler{}
	unpacker := Unpacker{}
	barcodeController := BarcodeController{}
	serialHandler := SerialHandler{}
	serialHandler.PortConfig("/dev/ttyAMA0", 9600)
	serialHandler.OpenPort()
	barcodeController.CreateBarcode(serialHandler.port)

	for {
		fmt.Print("Bar-code: ")
		barcodeOutput, err := barcodeController.barcode.ReadBytes('\x0d')
		if err != nil {
			panic(err)
		}
		stringBarcodeOutput := string(barcodeOutput)
		stringBarcodeOutput = strings.Replace(stringBarcodeOutput, "\r", "", -1)
		fmt.Println(stringBarcodeOutput)
		rawProductData := apiHandler.RequestProductData(stringBarcodeOutput)
		productData := unpacker.UnpackJSON(rawProductData)

		for key, value := range productData {
			fmt.Println(key, value)
		}
	}
}
