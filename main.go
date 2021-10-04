package main

import (
	"encoding/binary"
	"fmt"
	_ "github.com/lib/pq"
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
		barcodeAsByteArray := barcodeController.Read()


		i := int64(binary.LittleEndian.Uint64(barcodeAsByteArray))
		fmt.Println(i)


		/*stringBarcodeOutput := string(barcodeAsByteArray)
		stringBarcodeOutput = strings.Replace(stringBarcodeOutput, "\r", "", -1)*/

		/*var formatedBarcode int64
		if i, convErr := strconv.ParseInt(stringBarcodeOutput, 10, 64); convErr == nil {
			formatedBarcode = i
		}*/
		/*fmt.Println(formatedBarcode)
		postgreHandler.Connect()
		price, mj, mjkoef := postgreHandler.QueryProductData(formatedBarcode)
		postgreHandler.Disconnect()
		fmt.Println(price, mj, mjkoef)*/

	}
}