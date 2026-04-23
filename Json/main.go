package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	c := Cachorro{"Rex", "Dalmata", 3}

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]interface{}{
		"nome":  "toby",
		"raca":  "Poodle",
		"idade": 3,
	}
	c2JSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2JSON)
	fmt.Println(bytes.NewBuffer(c2JSON))

	aprendendoUnmarshal()

}
