package main

import "fmt"

func main()  {
	barcodeController := BarcodeController{}
	apiHandler := APIHandler{}
	unpacker := Unpacker{}

	barcodeController.CreateBarcode()

	for  {
		fmt.Print("Bar-code: ")
		productBarcode := barcodeController.ReadData()
		rawProductData := apiHandler.RequestProductData(productBarcode)
		productData := unpacker.UnpackJSON(rawProductData)

		for key, value := range productData {
			fmt.Println(key, value)
		}
	}
}
