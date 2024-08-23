package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Pruebas libreria net
	strUDP := ":1800"

	str, err := net.ResolveUDPAddr("udp4", strUDP)
	if err != nil {
		os.Exit(1)
	}

	dial, err := net.DialUDP("udp", nil, str)
	if err != nil {
		os.Exit(2)
	}

	dial.Write([]byte("Hola mundo desde el cliente"))

	var buffer [256]byte
	n, err := dial.Read(buffer[0:])
	if err != nil {
		os.Exit(3)
	}

	fmt.Println("Recibido", string(buffer[0:n]), "de", dial.RemoteAddr())

	dial.Close()

}
