package json

import (
	"encoding/json"
	"fmt"
)

func aprendendoUnmarshal() {
	cachorroEmJSON := `{"nome": "Rex", "raca": "Dalmata", "idade": 3}`

	var c Cachorro
	fmt.Println(c)

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		fmt.Println("Erro ao fazer unmarshal:", erro)
	}
	fmt.Println(c)

	cachorro2Em
}
