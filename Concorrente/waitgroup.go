package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done()
	}()
	go func() {
		escrever("Programação Concorrente")
		waitGroup.Done()
	}()
	//go escrever("Olá Mundo") //go inicia uma go routine, ou seja, uma thread leve, ou seja, uma função que pode ser executada concorrentemente com outras funções
	waitGroup.Wait() //espera as go routines terminarem
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
