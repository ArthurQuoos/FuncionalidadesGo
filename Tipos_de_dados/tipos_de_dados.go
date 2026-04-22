package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -10000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	numer := 123124567
	fmt.Println(numer)

	//alias
	//int32 é um alias para rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//Byte é um alias para uint8
	var numero4 byte = 255
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123456789.99
	fmt.Println(numeroReal2)

	numeroReal3 := 123456789.99
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'A' //Valor da tabela ASCII
	fmt.Println(char)

	texto := 5
	fmt.Println(texto)

	var booleano bool = true
	fmt.Println(booleano)

	var booleano2 bool
	fmt.Println(booleano2)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
