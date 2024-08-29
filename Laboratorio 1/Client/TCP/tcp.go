package tcp

import (
	"fmt"
	"net"
)

func ServerTCP(ip string, port int, numeroAleatorio int) (*net.TCPAddr, error) {
	// Resolver la dirección TCP del servidor
	address := fmt.Sprintf("%s:%d", ip, port)
	strTCP, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, fmt.Errorf("error al resolver la dirección TCP: %w", err)
	}
	fmt.Println("Dirección TCP resuelta:", strTCP)
	// Conectar al servidor TCP
	DialTCP, err := net.DialTCP("tcp", nil, strTCP)
	if err != nil {
		return nil, fmt.Errorf("error al conectar al servidor TCP: %w", err)
	}
	defer DialTCP.Close()
	// Convertir el número aleatorio a cadena y enviarlo al servidor
	mensajeTCP := fmt.Sprintf("%d", numeroAleatorio)
	_, err = DialTCP.Write([]byte(mensajeTCP))
	if err != nil {
		return nil, fmt.Errorf("error al enviar el mensaje: %w", err)
	}
	// Leer la respuesta del servidor
	var bufferTCP [256]byte
	nTCP, err := DialTCP.Read(bufferTCP[0:])
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta del servidor: %w", err)
	}
	fmt.Println("El servidor respondió en TCP:", string(bufferTCP[0:nTCP]))

	return strTCP, nil
}
