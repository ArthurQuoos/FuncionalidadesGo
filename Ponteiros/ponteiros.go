package main

import "fmt"

func main() {
	fmt.Println("Ponteiros em Golang")

	var variavel int = 10
	var ponteiro *int = &variavel

	fmt.Println("Valor de variavel:", variavel)
	fmt.Println("Endereço de variavel:", &variavel)
	fmt.Println("Valor de ponteiro (endereço de variavel):", ponteiro)
	fmt.Println("Valor apontado por ponteiro:", *ponteiro)

}
