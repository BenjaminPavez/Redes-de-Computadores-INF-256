package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
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
	reader := bufio.NewReader(DialTCP)

	for i := 0; i < numeroAleatorio; i++ {
		// Recibir la pregunta del servidor
		pregunta, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("error al recibir la pregunta: %w", err)
		}
		fmt.Print("Pregunta recibida: ", pregunta)

		// Recibir las alternativas del servidor
		for j := 0; j < 4; j++ {
			alternativa, err := reader.ReadString('\n')
			if err != nil {
				return nil, fmt.Errorf("error al recibir la alternativa: %w", err)
			}
			fmt.Printf("%s", alternativa)
		}

		// Mostrar la pregunta y las alternativas al usuario
		fmt.Print("Escribe tu respuesta (1-4): ")
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		// Convertir la respuesta del usuario a la posición correspondiente
		respuestaIndex, err := strconv.Atoi(userInput)
		if err != nil || respuestaIndex < 1 || respuestaIndex > 4 {
			fmt.Println("Respuesta inválida. Por favor, ingresa un número entre 1 y 4.")
			i-- // Repetir la misma pregunta
			continue
		}
		respuestaPosicion := respuestaIndex - 1

		// Enviar la respuesta del usuario al servidor
		_, err = DialTCP.Write([]byte(fmt.Sprintf("%d\n", respuestaPosicion)))
		if err != nil {
			return nil, fmt.Errorf("error al enviar la respuesta: %w", err)
		}

		// Recibir la verificación de la respuesta del servidor
		verificacion, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("error al recibir la verificación: %w", err)
		}
		fmt.Print("Verificación recibida: ", verificacion)
	}
	// Esperar el último mensaje del servidor (número de respuestas correctas)
	resultadoFinal, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error al recibir el resultado final: %w", err)
	}
	fmt.Print(resultadoFinal)
	return strTCP, nil
}
