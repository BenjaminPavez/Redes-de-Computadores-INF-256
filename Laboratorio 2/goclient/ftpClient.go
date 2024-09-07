package main

import (
	"os"
	"github.com/secsy/goftp"
)

const (
	ftpServerAddr = "192.168.1.188"
	ftpServerPath = "."       
)

//IronMoth
//ElectricoRoca
func main() {
	config := goftp.Config{
		User:     "grupo26",
		Password: "IronMoth",
	}
	client, err := goftp.DialConfig(config, ftpServerAddr)
	if err != nil {
		panic(err)
	}

	test, err := os.Create("Preguntas.txt")
	if err != nil {
		panic(err)

	}

	err = client.Retrieve("./Preguntas.txt", test)

	if err != nil {
		panic(err)
	}

}

