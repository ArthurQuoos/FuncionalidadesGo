package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func aprendendoUnmarshal() {
	cachorroEmJSON := `{"nome": "Rex", "raca": "Dalmata", "idade": 3}`

	var c Cachorro
	fmt.Println(c)

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorro2EmJSON := `{"nome": "Toby", "raca": "Poodle", "idade": 3}`
	c2 := make(map[string]interface{})
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
