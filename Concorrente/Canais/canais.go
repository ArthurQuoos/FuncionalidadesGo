package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo", canal)

	for {
		mensagem, aberto := <-canal //recebe o valor do canal, ou seja, o valor enviado para o canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal) //fecha o canal, PROTEJE CONTRA DEADLOCK, ou seja, contra o bloqueio de leitura do canal, ou seja, contra o bloqueio de escrita do canal
}
