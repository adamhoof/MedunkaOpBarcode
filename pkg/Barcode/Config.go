package barcode

import (
	"bufio"
	"github.com/tarm/serial"
)

func AssignPort(reader *bufio.Reader, port *serial.Port) {
	reader = bufio.NewReader(port)
}
