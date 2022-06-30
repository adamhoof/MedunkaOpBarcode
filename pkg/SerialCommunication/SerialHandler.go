package serialcommunication

import "github.com/tarm/serial"

func OpenPort(portName string, baud int) *serial.Port {
	port, err := serial.OpenPort(&serial.Config{Name: portName, Baud: baud})
	if err != nil {
		panic(err)
	}
	return port
}
