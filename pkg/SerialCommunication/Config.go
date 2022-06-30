package serialcommunication

import "github.com/tarm/serial"

func CreatePortConfig(portName string, baud int) serial.Config {
	return serial.Config{Name: portName, Baud: baud}
}
