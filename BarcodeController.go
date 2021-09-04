package main

import (
	"bufio"
	"os"
	"strings"
)

type BarcodeController struct {
	barcode *bufio.Reader
}

func (barcodeController *BarcodeController) CreateBarcode(){
	barcodeController.barcode = bufio.NewReader(os.Stdin)
}

func (barcodeController *BarcodeController) ReadData() string{
	name, _ := barcodeController.barcode.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	return name
}
