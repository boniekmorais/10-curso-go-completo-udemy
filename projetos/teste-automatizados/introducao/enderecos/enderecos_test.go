package enderecos

import (
	"testing"
)

func TestTipoDeEndereco(t *testing.T) {

	enderecoParaTeste := "Rua Paulista"
	tipoEsperado := "Avenida"

	tipoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoRecebido != tipoEsperado {
		// t.Error("O tipo recebido é diferente do esperado")
		t.Errorf("O tipo recebido é diferente do esperado. Esperado: %s | Recebido: %s", tipoEsperado, tipoRecebido)
	}
}
