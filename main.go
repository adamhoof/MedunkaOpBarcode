package main

import (
	_ "github.com/lib/pq"
)

func main() {
	postgreHandler := PostgreSQLHandler{}
	postgreHandler.Connect()
	/*postgreHandler.TestConnection()*/
	/*postgreHandler.CreateTable()*/
	postgreHandler.ImportFromCSV()
	postgreHandler.Disconnect()


	/*apiHandler := APIHandler{}
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

		fmt.Println()
		fmt.Println("Cena/Price:")
		fmt.Print(productData["price"])
		fmt.Println(",-")
		fmt.Println()
		fmt.Print("Jednotková cena/")
		fmt.Println("Unit Price:")
		fmt.Println(productData["price_in_sale"])*/
}
