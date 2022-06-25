package Barcode

import (
	"bufio"
	"github.com/tarm/serial"
)

type BarcodeController struct {
	barcode *bufio.Reader
}

func (barcodeController *BarcodeController) CreateBarcodeReader(port *serial.Port) {
	barcodeController.barcode = bufio.NewReader(port)
}

func (barcodeController *BarcodeController) Read() []byte {
	barcode, err := barcodeController.barcode.ReadBytes('\x0d')
	if err != nil {
		panic(err)
	}
	return barcode
}
