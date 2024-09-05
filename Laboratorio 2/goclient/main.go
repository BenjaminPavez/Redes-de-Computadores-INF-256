package main

import (
	"os"

	"github.com/secsy/goftp"
)

const (
	ftpServerAddr = "192.168.1.188" //actualizar si es necesario
	ftpServerPath = "."             //mantener esta línea para trabajar en la carpeta raíz del servidor
)

func main() {
	config := goftp.Config{
		User:     "grupo26",       //utilice el usuario asignado a su grupo
		Password: "ElectricoRoca", //utilice la contraseña obtenida durante la interacción con el servidor TCP
	}
	client, err := goftp.DialConfig(config, ftpServerAddr)
	if err != nil {
		panic(err)
	}

	test, err := os.Create("preguntas.txt") //utilice el nombre del archivo proporcionado en el laboratorio
	if err != nil {
		panic(err)

	}

	err = client.Retrieve("./preguntas.txt", test) //utilice el nombre del archivo proporcionado en el laboratorio

	if err != nil {
		panic(err)
	}

	// bigFile, err := os.Open("tcp.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// err = client.Store("./tcp.txt", bigFile)
	// if err != nil {
	// 	panic(err)
	// }
}
