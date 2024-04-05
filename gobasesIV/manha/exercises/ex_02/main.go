package main

import (
	"errors"
	"fmt"
)

/*
	Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de
	“Error()”, seja implementado “errors.New()”.
*/

func main() {
	salario := 12000

	if salario < 15000 {
		err := errors.New("erro: O salário digitado não está dentro do valor mínimo")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
