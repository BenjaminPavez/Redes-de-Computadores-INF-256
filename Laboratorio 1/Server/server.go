package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Questions struct {
    Questions []Question `json:"Question"`
}



type Question struct {
	Quest string `json:"Question"`
	Alts [4]string `json:"Alternatives"`
	Ans string `json:"Answer"`
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

	}

}



/*
La función se encarga de recibir mensajes del cliente y enviar mensajes al cliente

Parametros :

	numeroPreguntas : int - Número de preguntas que se le enviarán al cliente
	ip : string - Dirección IP del cliente
	puerto : int - Puerto del cliente

Retorno :

	Nada, no retorna ningun valor
*/
func ClientTCP(numeroPreguntas int, ip string, puerto int) {
    jsonFile, err := os.Open("test.json")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("El JSON 'test.json' se cargo correctamente")
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var question Questions

    json.Unmarshal(byteValue, &question)



	for i := 0; i < len(question.Questions); i++ {
        fmt.Println("Pregunta: " + question.Questions[i].Quest)
        fmt.Println("Alternativas: " + question.Questions[i].Alts[0] + ", " + question.Questions[i].Alts[1] + ", " + question.Questions[i].Alts[2] + ", " + question.Questions[i].Alts[3])
        fmt.Println("Respeusta: " + question.Questions[i].Ans)
    }

	
	// Escuchar en el puerto especificado
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", puerto))
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Servidor escuchando en el puerto %d\n", puerto)

	var flag bool = true
	for flag {
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
		numeroPreguntasRecibida, err := strconv.Atoi(numeroPreguntasStr)
		if err != nil {
			fmt.Println("Error al convertir el número de preguntas:", err)
			return
		}
		fmt.Printf("Número de preguntas recibido: %d\n", numeroPreguntasRecibida)
		if numeroPreguntasRecibida == 404 {
			fmt.Println("Cliente desconectado")
			listener.Close()
		}
		var respuestasCorrectas int = 0
		for {
			if numeroPreguntas == numeroPreguntasRecibida {
				for i := 0; i < numeroPreguntasRecibida; i++ {
					trivia := question.Questions[i%len(question.Questions)]
					pregunta := fmt.Sprintf("%s\n1. %s\n2. %s\n3. %s\n4. %s\n", trivia.Quest, trivia.Alts[0], trivia.Alts[1], trivia.Alts[2], trivia.Alts[3])
					_, err := conn.Write([]byte(pregunta))
					if err != nil {
						fmt.Println("Error al enviar la pregunta:", err)
						break
					}

					n, err := conn.Read(buffer)
					if err != nil {
						fmt.Println("Error al leer la respuesta del cliente:", err)
						break
					}
					respuestaCliente := strings.TrimSpace(string(buffer[:n]))
					fmt.Printf("Respuesta del cliente: %s\n", respuestaCliente)

					respuestaIndex, err := strconv.Atoi(respuestaCliente)
					if err != nil {
						fmt.Println("Error al convertir la respuesta del cliente a un índice:", err)
						_, err = conn.Write([]byte("Respuesta inválida\n"))
						if err != nil {
							fmt.Println("Error al enviar la respuesta inválida:", err)
						}
						return
					}
					if respuestaIndex >= 0 && respuestaIndex < len(trivia.Alts) && trivia.Alts[respuestaIndex] == trivia.Ans {
						fmt.Println("Respuesta correcta")
						respuestasCorrectas++
						_, err = conn.Write([]byte("Respuesta correcta\n"))
					} else {
						fmt.Println("Respuesta incorrecta")
						_, err = conn.Write([]byte("Respuesta incorrecta\n"))
					}
					if err != nil {
						fmt.Println("Error al enviar la verificación de la respuesta:", err)
						break
					}
				}
				resultado := fmt.Sprintf("Número de respuestas correctas: %d\n", respuestasCorrectas)
				_, err = conn.Write([]byte(resultado))
				if err != nil {
					fmt.Println("Error al enviar el resultado:", err)
				}
				break

			} else {
				break
			}
		}
		flag = false
		fmt.Println("Cliente desconectado")

	}
}



func main() {
	numeroPregunntas, ip, puerto := ClientUDP()
	fmt.Println("Numero de preguntas: ", numeroPregunntas)
	fmt.Println("IP: ", ip)
	fmt.Println("Puerto: ", puerto)
	ClientTCP(numeroPregunntas, ip, puerto)
}
