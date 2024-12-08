package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	serverAddr := net.UDPAddr{
		Port: 8080,                       //Puerto a conectarse
		IP:   net.ParseIP("192.168.1.3"), //Ip a conectarse
	}

	// Conectar al servidor UDP
	conn, err := net.DialUDP("udp", nil, &serverAddr)
	if err != nil {
		log.Fatalf("Error al conectarse al servidor UDP: %v", err)
	}
	defer conn.Close()

	ticker := time.NewTicker(2 * time.Second) //Establecer intervalo de envio del mensaje
	defer ticker.Stop()

	timeout := time.After(500 * time.Second) //Establecer Tiempo total de la conexion

	fmt.Println("Cliente UDP conectado. Enviando mensajes cada 5 segundos...")

	for {
		select {
		case <-timeout:
			fmt.Println("Tiempo cumplido. Cerrando cliente.")
			return
		case <-ticker.C:
			// Preparar el mensaje para enviar
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			message := fmt.Sprintf("hola, soy yo") //Ingrese mensaje que quiera enviar
			n, err := conn.Write([]byte(message))
			if err != nil {
				log.Printf("Error al enviar el mensaje: %v", err)
				return
			}
			fmt.Printf("[%s] Mensaje enviado: %s\nTamaño del paquete enviado: %d bytes\n", currentTime, message, n)

		}
	}
}
