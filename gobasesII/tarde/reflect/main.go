package main

import (
	"fmt"
	"reflect"
)

type Pessoa struct {
	PrimeiroNome string `bd:"primeiro_nome"`
	Sobrenome    string `bd:"sobrenome"`
	Telefone     string `bd:"telefone"`
	Endereco     string `bd:"endereco"`
}

func main() {
	pessoa := Pessoa{}
	p := reflect.TypeOf(pessoa)

	fmt.Println(p)

	fmt.Println("Type: ", p.Name())
	fmt.Println("Kind: ", p.Kind())

	// não podemos usar o range
	for i := 0; i < p.NumField(); i++ {
		field := p.Field(i)
		name := field.Name
		tag := field.Tag.Get("bd")
		fmt.Printf("field: %s, tag: %s \n", name, tag)
	}

}
