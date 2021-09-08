package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"time"
)

func main() {
	/*apiHandler := APIHandler{}
	unpacker := Unpacker{}*/

	config := &serial.Config{
		Name: "/dev/ttyAMA0",
		Baud: 9600,
		ReadTimeout: 1*time.Second,
	}
	stream, err := serial.OpenPort(config)
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
