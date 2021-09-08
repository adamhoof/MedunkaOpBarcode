package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
)

func main() {
	/*apiHandler := APIHandler{}
	unpacker := Unpacker{}*/
	barcodeController := BarcodeController{}
	barcodeController.CreateBarcode()
	barcodeController.SetupBarcode("/dev/ttyAMA0", 9600)
	/*barcodeController.barcode.ReadBytes('\n')*/

	stream, err := serial.OpenPort(barcodeController.config)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)

	for {
		fmt.Print("Bar-code: ")
		n, err := stream.Read(buf)
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
