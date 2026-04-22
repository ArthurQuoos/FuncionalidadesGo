package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices em GoLang")

	var array [5]string
	array[0] = "Posicao 1"

	fmt.Println("Array:", array)

	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println("Array 2:", array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array 3:", array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println("Slice 2:", slice2)

	array2[1] = "posicao alterada"
	fmt.Println("Array 2 alterado:", array2)

	//arrays internos
	fmt.Println("---------------------------")
	// make( tipo, tamanho, capacidade )
	slice3 := make([]float32, 10, 11)
	fmt.Println("Slice 3:", slice3)
	fmt.Println("Tamanho:", len(slice3))
	fmt.Println("Capacidade:", cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println("Slice 3 após append:", slice3)
	fmt.Println("Tamanho após append:", len(slice3))
	fmt.Println("Capacidade após append:", cap(slice3))

	slice4 := make([]float32, 0, 5)
	fmt.Println("Slice 4:", slice4)
	fmt.Println("Tamanho Slice 4:", len(slice4))
	fmt.Println("Capacidade Slice 4:", cap(slice4))
}
