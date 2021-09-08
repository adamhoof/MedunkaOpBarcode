package main

import (
	"fmt"
	"log"
)

func main() {
	/*apiHandler := APIHandler{}
	unpacker := Unpacker{}*/
	barcodeController := BarcodeController{}
	serialHandler := SerialHandler{}
	barcodeController.CreateBarcode()
	serialHandler.SetupBarcode("/dev/ttyAMA0", 9600)
	serialHandler.OpenPort()

	buf := make([]byte, 1024)

	for {
		fmt.Print("Bar-code: ")
		n, err := serialHandler.port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf[:n])
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
