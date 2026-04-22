package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func calculos(a, b int8) (int8, int8) {
	return a + b, a - b
}

func main() {
	soma := somar(5, 7)
	fmt.Println(soma)
	fmt.Println(somar(10, 20))

	var f = func(x int, y int) int {
		return x * y
	}

	resultado := f(3, 4)
	fmt.Println(resultado)

	soma2, subtracao := calculos(20, 10)
	fmt.Println(soma2, subtracao)

	soma3, _ := calculos(30, 20)
	fmt.Println(soma3)
}
