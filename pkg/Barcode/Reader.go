package barcode

import (
	"bufio"
)

func Read(reader *bufio.Reader) []byte {
	barcode, err := reader.ReadBytes('\x0d')
	if err != nil {
		panic(err)
	}
	return barcode
}
