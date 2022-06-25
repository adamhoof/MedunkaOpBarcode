package Barcode

import (
	"bufio"
	"github.com/tarm/serial"
)

func AssignPort(reader *bufio.Reader, port *serial.Port) {
	reader = bufio.NewReader(port)
}

func Read(reader *bufio.Reader) []byte {
	barcode, err := reader.ReadBytes('\x0d')
	if err != nil {
		panic(err)
	}
	return barcode
}
