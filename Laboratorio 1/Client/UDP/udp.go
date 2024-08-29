package udp

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
La función se encarga de limpiar la pantalla

Parametros :

	Nada, no recibe ningun parametro

Retorno :

	Nada, no retorna ningun valor
*/
func clearScreen() {
	fmt.Print("\033[H\033[2J") // Limpiar la pantalla
}

/*
La función se encarga de dar formato a la pregunta y sus alternativas

Parametros :

	[]byte : buffer, arreglo de bytes que contiene la pregunta y sus alternativas

Retorno :

	Nada, no retorna ningun valor
*/
func formato(buffer []byte) {
	data := string(buffer[0:])
	lines := strings.Split(data, "\n")
	if len(lines) < 5 {
		fmt.Println("Faltan datos.")
		return
	}
	pregunta := lines[0]
	alternativas := lines[1:5]
	fmt.Println("Pregunta:", pregunta)
	for i, alt := range alternativas {
		fmt.Printf(" %d. %s\n", i+1, alt)
	}
}

/*
La función se encarga de enviar y recibir mensajes del servidor

Parametros :

	int : opcion, opción seleccionada por el usuario

Retorno :

	Nada, no retorna ningun valor
*/
func ServerUDP(opcion int) (string, int, int) {
	strUDP := ":1800"

	str, err := net.ResolveUDPAddr("udp4", strUDP)
	if err != nil {
		fmt.Println("Error al resolver la dirección UDP:", err)
		os.Exit(1)
	}

	dial, err := net.DialUDP("udp", nil, str)
	if err != nil {
		fmt.Println("Error al conectar al servidor UDP:", err)
		os.Exit(2)
	}
	defer dial.Close()

	//Enviar la opción al servidor
	mensaje := fmt.Sprintf("%d", opcion)
	_, err = dial.Write([]byte(mensaje))
	if err != nil {
		fmt.Println("Error al enviar el mensaje:", err)
		os.Exit(4)
	}

	//Leer la respuesta del servidor (número de preguntas)
	var buffer [256]byte
	n, err := dial.Read(buffer[0:])
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor:", err)
		os.Exit(5)
	}
	// Convertir la respuesta a una cadena
	response := string(buffer[:n])

	// Analizar la respuesta del servidor
	// Expresión regular para extraer los valores
	re := regexp.MustCompile(`IP:\s*([0-9.]+),\s*Puerto:\s*(\d+),\s*Número Aleatorio:\s*(\d+)`)
	matches := re.FindStringSubmatch(response)

	if len(matches) != 4 {
		fmt.Println("Error al analizar la respuesta del servidor: formato no coincide")
		os.Exit(6)
	}
	serverIP := matches[1]
	serverPort, err := strconv.Atoi(matches[2])
	if err != nil {
		fmt.Println("Error al convertir el puerto del servidor:", err)
		os.Exit(7)
	}
	numeroAleatorio, err := strconv.Atoi(matches[3])
	if err != nil {
		fmt.Println("Error al convertir el número aleatorio:", err)
		os.Exit(8)
	}
	clearScreen()
	//Mostrar el número de preguntas
	fmt.Println("Tendras que responder:", numeroAleatorio, "preguntas")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("-------------------------------------El servidor respondio--------------------------------------")
	fmt.Println("-----------------------------------------Desea comenzar?----------------------------------------")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("1. Si")
	fmt.Println("2. No")
	fmt.Print("-> ")

	//Leer la respuesta del usuario
	var opt2 int
	_, err = fmt.Scanln(&opt2)
	if err != nil {
		fmt.Println("Error al leer la entrada del usuario:", err)
		os.Exit(6)
	}

	if opt2 == 1 {
		return serverIP, serverPort, numeroAleatorio
	}
	// En caso de que el usuario elija no comenzar, se debe retornar algo
	return "", 0, 0

}
