package maps

import (
	"fmt"
)

func Maps() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Pedro",
			"segundo":  "Silva",
		},
		"curso": {
			"nome":  "Engenharia",
			"Local": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "curso")

	fmt.Println(usuario2)
}
