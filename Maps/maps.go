package main

import "fmt"

func main() {
	fmt.Println("Maps em GoLang")

	usuario := map[string]string{
		"nome":      "Maria",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println("Nome:", usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Souza",
		},
		"endereco": {
			"logradouro": "Rua XYZ",
			"numero":     "456",
		},
	}
	fmt.Println("Usuário 2:", usuario2)
	fmt.Println("\nNome do usuário 2:", usuario2["nome"]["primeiro"], usuario2["nome"]["ultimo"])

	delete(usuario, "endereco") //remove a chave endereco do map usuario
	fmt.Println("\nUsuário após delete:", usuario)

	usuario2["signo"] = map[string]string{
		"tipo": "Leão",
	}
	fmt.Println("Usuário 2 após adicionar signo:", usuario2)

}
