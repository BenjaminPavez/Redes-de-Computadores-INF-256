package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// Crear una dirección UDP en el puerto 8080
	addr := net.UDPAddr{
		Port: 8080,                     //Puerto a escuchar
		IP:   net.IPv4(192, 168, 1, 3), //Ip del servidor
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor UDP: %v", err)
	}
	defer conn.Close()

	fmt.Println("Servidor UDP escuchando en el puerto 8080...")

	buffer := make([]byte, 1024)

	for {
		// Leer datos del cliente
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error al leer del cliente: %v", err)
			continue
		}

		currentTime := time.Now().Format("2006-01-02 15:04:05")
		message := string(buffer[:n])
		fmt.Printf("[%s] Mensaje recibido de %s: %sTamaño del paquete: %d bytes\n", currentTime, clientAddr, message, n)

		response := fmt.Sprintf("[%s] Mensaje recibido con éxito\n", currentTime)
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			log.Printf("Error al responder al cliente: %v", err)
		}
	}
}
