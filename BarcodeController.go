package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
)

type BarcodeController struct {
	barcode *bufio.Reader
}

func (barcodeController *BarcodeController) CreateBarcode(port *serial.Port) {
	barcodeController.barcode = bufio.NewReader(port)
}

func (barcodeController *BarcodeController) ReadData() {
	code, _ := barcodeController.barcode.ReadString('\n')
	fmt.Println(code)
}