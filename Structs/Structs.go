package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}
type usuario struct {
	pessoa   //"Heranca" de pessoa
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs em GoLang")

	var u usuario
	u.nome = "João"
	u.idade = 25

	fmt.Println("Nome do usuário:", u.nome)
	fmt.Println("Idade do usuário:", u.idade)

	enderecoExemplo := endereco{"Rua ABC", 123}
	u.endereco = enderecoExemplo

	fmt.Println("Endereço do usuário:", u.endereco.logradouro, ",", u.endereco.numero)
}
