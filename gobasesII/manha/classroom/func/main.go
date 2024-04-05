package main

import "fmt"

func main() {
	a, b, c, d := 1, 0, 5, -3

	verificarVariavel(a)
	verificarVariavel(b)
	verificarVariavel(c)
	verificarVariavel(d)
}

func verificarVariavel(numero int) {
	if numero < 0 {
		fmt.Println("O número é negativo")
	} else if numero > 0 {
		fmt.Println("O número é positivo")
	} else {
		fmt.Println("O número é zero")
	}
}

func minhaFuncao(valor1, valor2 float64) {}
