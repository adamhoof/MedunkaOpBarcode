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
		barcodeOutput, err := barcodeController.barcode.ReadBytes('\x0d')
		if err != nil {
			panic(err)
		}
		fmt.Print("\033[H\033[2J")
		stringBarcodeOutput := string(barcodeOutput)
		stringBarcodeOutput = strings.Replace(stringBarcodeOutput, "\r", "", -1)
		rawProductData := apiHandler.RequestProductData(stringBarcodeOutput)
		productData := unpacker.UnpackJSON(rawProductData)

		fmt.Println("Cena/Price: ")
		fmt.Println(productData["price"])
		fmt.Println("Jednotková cena: ")
		fmt.Println("Unit Price: ")
		fmt.Println(productData["price_in_sale"])
	}
}
