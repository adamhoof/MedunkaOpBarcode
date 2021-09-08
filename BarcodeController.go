package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
)

type BarcodeController struct {
	barcode *bufio.Reader
	config serial.Config
}

func (barcodeController *BarcodeController) CreateBarcode() {
	barcodeController.barcode = bufio.NewReader(&serial.Port{})
}

func (barcodeController *BarcodeController) SetupBarcode(port string, baud int) {
	barcodeController.config.Name = port
	barcodeController.config.Baud = baud
}

func (barcodeController *BarcodeController) ReadData() {
	code, _ := barcodeController.barcode.ReadString('\n')
	fmt.Println(code)
}