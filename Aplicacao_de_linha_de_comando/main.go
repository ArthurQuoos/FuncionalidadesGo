package main

import (
	"fmt"
	"linha_de_comando/app"
	"log"
	"os"
)

func main() {

	fmt.Println("Aplicacao de linha de comando em GoLang")

	aplicacao := app.Gerar()
	erroaplicacao := aplicacao.Run(os.Args)
	if erroaplicacao != nil {
		log.Fatal(erroaplicacao)
	}

}
