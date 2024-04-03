package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	PrimeiroNome string `json:"primeiroNome" bd:"primeiro_nome"`
	Sobrenome    string `json:"sobrenome" bd:"sobrenome"`
	Telefone     string `json:"telefone" bd:"telefone"`
	Endereco     string `json:"endereco" bd:"endereco"`
}

func main() {
	p1 := Pessoa{
		PrimeiroNome: "Fulano",
		Sobrenome:    "Cliclano",
		Telefone:     "123456789",
		Endereco:     "Rua dos bobos, n° 0",
	}

	fmt.Printf("%+v \n", p1)

	json, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("Algum erro aconteceu ao parsear o json", err)
	}

	fmt.Println(string(json))
	fmt.Println(err)
}
