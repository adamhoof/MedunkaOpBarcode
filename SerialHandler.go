package main

import "github.com/tarm/serial"

type SerialHandler struct {
	config serial.Config
	port *serial.Port
}

func (serialHandler *SerialHandler) PortConfig(port string, baud int) {
	serialHandler.config.Name = port
	serialHandler.config.Baud = baud
}

func (serialHandler *SerialHandler) OpenPort(){
	port, err := serial.OpenPort(&serialHandler.config)
	if err != nil {
		panic(err)
	}
	serialHandler.port = port
}
