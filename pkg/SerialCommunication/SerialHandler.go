package SerialCommunication

import "github.com/tarm/serial"

func CreatePortConfig(portName string, baud int) serial.Config {
	return serial.Config{Name: portName, Baud: baud}
}

func OpenPort(config serial.Config) *serial.Port {
	port, err := serial.OpenPort(&config)
	if err != nil {
		panic(err)
	}
	return port
}
