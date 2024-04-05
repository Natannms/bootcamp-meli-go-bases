package main

import "fmt"

/*
Faça uma aplicação que contenha uma variável com o número do mês.
Dependendo do número, imprima o mês correspondente em texto.
Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você escolheria e porquê?
Ex: 7 de julho.
*/

func main() {
	meses := []string{
		"Janeiro",
		"Fevereiro",
		"Março",
		"Abril",
		"Maio",
		"Junho",
		"Julho",
		"Agosto",
		"Setembro",
		"Outubro",
		"Novembro",
		"Dezembro",
	}

	mes1 := 1   //Janeiro
	mes10 := 10 //Outubro
	mes12 := 12 //Dezembro

	fmt.Println("mês 1:", meses[mes1-1])
	fmt.Println("mês 10:", meses[mes10-1])
	fmt.Println("mês 12:", meses[mes12-1])

}
