package main

import "fmt"

/*
Um estudante de programação tentou fazer declarações de variáveis ​​com seus respectivos tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e pediu a ajuda de um desenvolvedor experiente que pode:
Revisar o código e realizar as devidas correções.
*/

func main() {
	// before
	// var sobrenome string = "Silva"
	// var idade int = "25"
	// boolean := "false";
	// var salario string = 4585.90
	// var nome string = "Fellipe"

	// after
	var nome string = "Fellipe"
	var sobrenome string = "Silva"
	var idade int = 25
	boolean := false
	var salario float64 = 4585.90

	fmt.Println(nome, sobrenome, idade, boolean, salario)

}
