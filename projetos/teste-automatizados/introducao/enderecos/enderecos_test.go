package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel() // Roda testes em paralelo.

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Bandeirantes", "Rodovia"},
	}

	for _, cenario := range cenariosDeTeste {

		enderecoParaTeste := cenario.enderecoInserido
		tipoEsperado := cenario.retornoEsperado

		tipoRecebido := TipoDeEndereco(enderecoParaTeste)

		if tipoRecebido != tipoEsperado {
			t.Error("O tipo recebido é diferente do esperado")
			t.Errorf("O tipo recebido é diferente do esperado. Esperado: %s | Recebido: %s", tipoEsperado, tipoRecebido)
		}
	}

	// enderecoParaTeste := "Rua Paulista"
	// tipoEsperado := "Avenida"

	// tipoRecebido := TipoDeEndereco(enderecoParaTeste)

	// if tipoRecebido != tipoEsperado {
	// t.Error("O tipo recebido é diferente do esperado")
	// t.Errorf("O tipo recebido é diferente do esperado. Esperado: %s | Recebido: %s", tipoEsperado, tipoRecebido)
	// }
}

// go test -v ./...
// go test --cover ./...
// go test --coverprofile cobertura.txt
// go tool cover --func=cobertura.txt
// go tool cover --html=cobertura.txt
