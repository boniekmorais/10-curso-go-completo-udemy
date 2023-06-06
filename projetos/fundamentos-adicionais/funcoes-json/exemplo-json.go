package funcoesjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

// type cachorro struct {
// 	Nome  string `json:"."`      Ignora a propriedade
// 	Raca  string `json:"raca"`
// 	Idade uint   `json:"idade"`
// }

func ExemploJsonMarshal() {

	c := cachorro{"Rex", "Dalmata", 3}

	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome":  "Toby",
		"raca":  "Poodle",
		"idade": "5",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}

func ExemploJsonUnmarshal() {

	cachorroEmJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`
	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Rexel","raca":"Dalmata","idade":"5"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
