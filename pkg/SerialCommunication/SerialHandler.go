package serialcommunication

import "github.com/tarm/serial"

func OpenPort(config *serial.Config) *serial.Port {
	port, err := serial.OpenPort(config)
	if err != nil {
		panic(err)
	}
	return port
}
