package main

import (
	"fmt"
	"time"
)

func opSoma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func opSubtra(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func opMultip(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

func opDivis(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}

func operacaoAritmetica(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case "Soma":
		return opSoma
	case "Subtracao":
		return opSubtra
	case "Multip":
		return opMultip
	case "Divis":
		return opDivis
	}

	return nil
}

func main() {
	oper := operacaoAritmetica("Subtracao")
	result := oper(2, 5)
	fmt.Println(result)

	// debounce
	time.AfterFunc(time.Millisecond*300, func() {
		fmt.Println("Alguma coisa")
	})

	time.Sleep(time.Second)
}
