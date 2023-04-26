package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEndereco: Verifica se o endereço tem um tipo válido e o retorna.

func TipoDeEndereco(endereco string) string {

	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoMinusculo, " ")[0]

	enderecoTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTipoValido = true
		}
	}

	if enderecoTipoValido {

		// Deprecated
		// return strings.Title(primeiraPalavra)

		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(primeiraPalavra)
	}

	return "Tipo inválido"

}
