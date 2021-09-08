package main

import (
	"fmt"
)

func main() {
	/*apiHandler := APIHandler{}
	unpacker := Unpacker{}*/
	barcodeController := BarcodeController{}
	serialHandler := SerialHandler{}
	serialHandler.PortConfig("/dev/ttyAMA0", 9600)
	serialHandler.OpenPort()
	barcodeController.CreateBarcode(serialHandler.port)


	for {
		fmt.Print("Bar-code: ")
		n, err := barcodeController.barcode.ReadBytes('\x0d')
		s := string(n)
		if err != nil {
			panic(err)
		}
		fmt.Println(s)
		/*rawProductData := apiHandler.RequestProductData(s)
		productData := unpacker.UnpackJSON(rawProductData)
		for key, value := range productData {
			fmt.Println(key, value)
		}*/
		/*for  {
		fmt.Print("Bar-code: ")
		productBarcode := barcodeController.ReadData()
		rawProductData := apiHandler.RequestProductData(productBarcode)
		productData := unpacker.UnpackJSON(rawProductData)

		for key, value := range productData {
			fmt.Println(key, value)
		}*/
	}
}
