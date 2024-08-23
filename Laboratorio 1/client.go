package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	strUDP := ":1800"

	// Resolver la dirección UDP
	str, err := net.ResolveUDPAddr("udp4", strUDP)
	if err != nil {
		fmt.Println("Error al resolver la dirección UDP:", err)
		os.Exit(1)
	}

	// Conectar al servidor UDP
	dial, err := net.DialUDP("udp", nil, str)
	if err != nil {
		fmt.Println("Error al conectar al servidor UDP:", err)
		os.Exit(2)
	}
	defer dial.Close()

	// Leer el mensaje del usuario
	var mensaje string
	fmt.Print("Escriba el mensaje: ")
	_, err = fmt.Scanln(&mensaje)
	if err != nil {
		fmt.Println("Error al leer el mensaje:", err)
		os.Exit(3)
	}

	// Enviar el mensaje al servidor
	_, err = dial.Write([]byte(mensaje))
	if err != nil {
		fmt.Println("Error al enviar el mensaje:", err)
		os.Exit(4)
	}

	// Leer la respuesta del servidor
	var buffer [256]byte
	n, err := dial.Read(buffer[0:])
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor:", err)
		os.Exit(5)
	}

	fmt.Println("El servidor respondió:", string(buffer[0:n]))
}
