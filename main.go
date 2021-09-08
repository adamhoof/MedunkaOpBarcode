package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	/*apiHandler := APIHandler{}
	unpacker := Unpacker{}*/
	barcodeController := BarcodeController{}
	serialHandler := SerialHandler{}
	barcodeController.CreateBarcode()
	serialHandler.PortConfig("/dev/ttyAMA0", 9600)
	serialHandler.OpenPort()

	for {
		fmt.Print("Bar-code: ")
		n, err := barcodeController.barcode.ReadString('\n')
		n = strings.Replace(n, "\n", "", -1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n)
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
