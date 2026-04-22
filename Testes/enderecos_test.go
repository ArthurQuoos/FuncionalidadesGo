// Teste de Unidade
package Testes

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida XYZ", "Avenida"},
		{"Estrada 123", "Estrada"},
		{"Rodovia 456", "Rodovia"},
		{"Praça Central", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tiposDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tiposDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereço recebido é diferente do esperado. Recebido: %s, Esperado: %s",
				tiposDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}
