package main

import (
	"fmt"
	tcp "main/Client/TCP"
	udp "main/Client/UDP"
	"os"
)

/*
La función se encarga de mostrar el menú al usuario

Parametros :

	Nada, no recibe ningun parametro

Retorno :

	Nada, no retorna ningun valor
*/
func MenuClient() {
	var opt int
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------Bienvenido a la Trivia USM-----------------------------------")
	fmt.Println("----------------------------------------Que desea hacer?----------------------------------------")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("1. Comenzar la ronda de preguntas")
	fmt.Println("2. Salir")
	fmt.Print("-> ")
	_, err := fmt.Scanln(&opt)
	if err != nil {
		fmt.Println("Error al leer la entrada del usuario:", err)
		os.Exit(1)
	}

	if opt == 1 {
		serverIP, serverPort, numeroAleatorio := udp.ServerUDP(opt)
		// Uso de los valores retornados
		fmt.Printf("IP del servidor: %s\n", serverIP)
		fmt.Printf("Puerto del servidor: %d\n", serverPort)
		fmt.Printf("Número aleatorio: %d\n", numeroAleatorio)
		tcp.ServerTCP(serverIP, serverPort, numeroAleatorio)

	} else if opt == 2 {
		os.Exit(0)
	} else {
		MenuClient()
	}
}

func main() {
	MenuClient()

	/*
		var err error
		for i := 0; i <  2; i++ {
			fmt.Print("Escriba el mensaje: ")
			_, err = fmt.Scanln(&pregunta)
			if err != nil {
				fmt.Println("Error al leer el mensaje:", err)
				return
			}
			strTCP, err := tcp.EnviarMensajeTCP(pregunta)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Mensaje enviado a la dirección TCP:", strTCP)

		}
	*/
}
