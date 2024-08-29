package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

type Trivia struct {
	Question     string
	Alternatives [4]string
	Answer       string
}

// Array de preguntas y respuestas
var QA = []Trivia{
	{
		Question:     "¿Cuál es el planeta más cercano al Sol?",
		Alternatives: [4]string{"Marte", "Venus", "Mercurio", "Júpiter"},
		Answer:       "Mercurio",
	},
	{
		Question:     "¿En qué continente se encuentra el desierto del Sahara?",
		Alternatives: [4]string{"Asia", "América", "Australia", "África"},
		Answer:       "África",
	},
	{
		Question:     "¿Cuál es el idioma más hablado en el mundo?",
		Alternatives: [4]string{"Inglés", "Español", "Chino mandarín", "Árabe"},
		Answer:       "Chino mandarín",
	},
	{
		Question:     "¿Cuál es el océano más grande del mundo?",
		Alternatives: [4]string{"Océano Atlántico", "Océano Índico", "Océano Pacífico", "Océano Ártico"},
		Answer:       "Océano Pacífico",
	},
	{
		Question:     "¿Qué elemento químico tiene como símbolo H?",
		Alternatives: [4]string{"Helio", "Hidrógeno", "Hierro", "Carbono"},
		Answer:       "Hidrógeno",
	},
	{
		Question:     "¿Quién pintó la Mona Lisa?",
		Alternatives: [4]string{"Pablo Picasso", "Vincent van Gogh", "Leonardo da Vinci", "Miguel Ángel"},
		Answer:       "Leonardo da Vinci",
	},
	{
		Question:     "¿Cuál es el animal terrestre más rápido?",
		Alternatives: [4]string{"Tigre", "Guepardo", "León", "Elefante"},
		Answer:       "Guepardo",
	},
	{
		Question:     "¿En qué año llegó el hombre a la Luna por primera vez?",
		Alternatives: [4]string{"1965", "1969", "1972", "1959"},
		Answer:       "1969",
	},
	{
		Question:     "¿Cuál es el río más largo del mundo?",
		Alternatives: [4]string{"Nilo", "Amazonas", "Yangtsé", "Misisipi"},
		Answer:       "Amazonas",
	},
	{
		Question:     "¿Qué instrumento musical tiene teclas negras y blancas?",
		Alternatives: [4]string{"Guitarra", "Violín", "Piano", "Flauta"},
		Answer:       "Piano",
	},
}

// Función para escribir en el servidor TCP.
func writeToTCP(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error al enviar mensaje TCP:", err)
	}
	conn.Close()
}

/*
La función se encarga de recibir mensajes del cliente y enviar mensajes al cliente
Parametros :

	Nada, no recibe ningun parametro

Retorno :

	Nada, no retorna ningun valor
*/
func ClientUDP() (int, string, int) {
	strUDP := ":1800"
	str, err := net.ResolveUDPAddr("udp4", strUDP)
	if err != nil {
		fmt.Println("Error al resolver la dirección UDP:", err)
		os.Exit(1)
	}
	udp, err := net.ListenUDP("udp", str)
	if err != nil {
		fmt.Println("Error al escuchar en UDP:", err)
		os.Exit(2)
	}
	defer udp.Close()
	fmt.Println("Server On")
	// Inicio del Server.
	var numeroAleatorio int = 0
	for {
		if numeroAleatorio == 0 {
			rand.Seed(time.Now().UnixNano())
			numeroAleatorio = rand.Intn(5) + 3
			fmt.Println("Número aleatorio generado:", numeroAleatorio)
		}
		var message [1024]byte
		n, addr, err := udp.ReadFromUDP(message[0:])
		if err != nil {
			fmt.Println("Error al leer el mensaje:", err)
			continue
		}
		fmt.Println("Mensaje recibido:", string(message[0:n]))
		//Obtener la dirección IP y el puerto del cliente
		clientIP := addr.IP.String()
		clientPort := addr.Port
		fmt.Printf("Dirección del cliente: %s:%d\n", clientIP, clientPort)

		// Enviar los datos de conexión y el número aleatorio de vuelta al cliente
		response := fmt.Sprintf("IP: %s, Puerto: %d, Número Aleatorio: %d", clientIP, clientPort, numeroAleatorio)
		_, err = udp.WriteToUDP([]byte(response), addr)
		if err != nil {
			fmt.Println("Error al enviar los datos de conexión:", err)
		}
		return numeroAleatorio, clientIP, clientPort
		/*
			strMessage := string(message[0:n])
			intMessage, err := strconv.Atoi(strMessage)
			if err != nil {
				fmt.Println("Error al convertir el mensaje:", err)
				continue
			}
			if intMessage == 1 {
				strRespt := fmt.Sprintf("%d", numeroAleatorio)
				_, err = udp.WriteToUDP([]byte(strRespt), addr)
				if err != nil {
					fmt.Println("Error al enviar el mensaje:", err)
				}
			} else if intMessage == 2 {
				preguntaAl := rand.Intn(len(QA))
				//Formatear el mensaje
				trivia := QA[preguntaAl]
				fmt.Println("Pregunta para el usuario: ", trivia)
				strRespt := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", trivia.Question, trivia.Alternatives[0], trivia.Alternatives[1], trivia.Alternatives[2], trivia.Alternatives[3])
				_, err = udp.WriteToUDP([]byte(strRespt), addr)
				if err != nil {
					fmt.Println("Error al enviar el mensaje:", err)
				}
				// Leer la respuesta del cliente
				n, _, err := udp.ReadFromUDP(message[0:])
				if err != nil {
					fmt.Println("Error al leer la respuesta del cliente:", err)
					continue
				}
				respuestaCliente := string(message[0:n])
				intResp, _ := strconv.Atoi(respuestaCliente)
				fmt.Println("El usuario responde: ", trivia.Alternatives[intResp-1])
				//Verificar si la respuesta es correcta
				if trivia.Alternatives[intResp-1] == trivia.Answer {
					fmt.Println("El usuario respondio correctamente")
					_, err = udp.WriteToUDP([]byte("correcta"), addr)
					if err != nil {
						fmt.Println("Error al enviar el mensaje:", err)
					}
				} else {
					fmt.Println("El usuario respondio incorrectamente")
					_, err = udp.WriteToUDP([]byte("incorrecta"), addr)
					if err != nil {
						fmt.Println("Error al enviar el mensaje:", err)
					}
				}
			|*/
	}

}

func ClientTCP(numeroPreguntas int, ip string, puerto int) {
	// Escuchar en el puerto especificado
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", puerto))
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Servidor escuchando en el puerto %d\n", puerto)

	for {
		// Aceptar una nueva conexión
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		fmt.Println("Cliente conectado")

		// Manejar la conexión
		defer conn.Close()

		// Leer el número de preguntas del cliente
		buffer := make([]byte, 256)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer el número de preguntas:", err)
			return
		}
		numeroPreguntasStr := string(buffer[:n])
		numeroPreguntas, err := strconv.Atoi(numeroPreguntasStr)
		if err != nil {
			fmt.Println("Error al convertir el número de preguntas:", err)
			return
		}
		fmt.Printf("Número de preguntas recibido: %d\n", numeroPreguntas)

		for {
			// Leer el mensaje del cliente
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error al leer el mensaje del cliente:", err)
				break
			}
			message := string(buffer[:n])
			fmt.Print("Mensaje del cliente: ", message)
		}
	}
}

func main() {
	numeroPregunntas, ip, puerto := ClientUDP()
	fmt.Println("Numero de preguntas: ", numeroPregunntas)
	fmt.Println("IP: ", ip)
	fmt.Println("Puerto: ", puerto)
	ClientTCP(numeroPregunntas, ip, puerto)
}
