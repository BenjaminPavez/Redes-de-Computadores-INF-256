package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*
La función se encarga de conectarse al servidor TCP y enviarle un número aleatorio

Parametros :

	ip : string - Dirección IP del servidor
	port : int - Puerto del servidor
	numeroAleatorio : int - Número aleatorio que se le enviará al servidor

Retorno :

	*net.TCPAddr - Dirección TCP del servidor
	error - Error al resolver la dirección TCP o al conectar al servidor TCP
*/
func ServerTCP(ip string, port int) (*net.TCPAddr, error) {
	//Resolver la dirección TCP del servidor
	address := fmt.Sprintf("%s:%d", ip, port)
	strTCP, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, fmt.Errorf("error al resolver la dirección TCP: %w", err)
	}
	fmt.Println("Dirección TCP resuelta:", strTCP)

	DialTCP, err := net.DialTCP("tcp", nil, strTCP)
	if err != nil {
		return nil, fmt.Errorf("error al conectar al servidor TCP: %w", err)
	}

	reader := bufio.NewReader(DialTCP)
	pregunta, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error al recibir la pregunta: %w", err)
	}
	fmt.Print("Resp Servidor: ", pregunta)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	id := scanner.Text()
	fmt.Fprintln(DialTCP, id)

	pregunta2, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error al recibir la pregunta: %w", err)
	}
	fmt.Print("Resp Servidor: ", pregunta2)
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	id1 := scanner1.Text()
	fmt.Fprintln(DialTCP, id1)

	pregunta3, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error al recibir la pregunta: %w", err)
	}
	fmt.Print("Resp Servidor: ", pregunta3)
	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	id3 := scanner3.Text()
	fmt.Fprintln(DialTCP, id3)

	DialTCP.Close()
	return strTCP, nil
}

func main() {
	ServerTCP("192.168.1.188", 8080)
}
