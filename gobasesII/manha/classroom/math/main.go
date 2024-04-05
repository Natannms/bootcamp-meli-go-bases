package main

import "fmt"

const (
	Soma   = "+"
	Subtra = "-"
	Multip = "*"
	Divis  = "/"
)

func soma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func operacaoAritmetica(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case Soma:
		return valor1 + valor2
	case Subtra:
		return valor1 - valor2
	case Multip:
		return valor1 * valor2
	case Divis:
		if valor2 != 0 {
			return valor1 / valor2
		}
	}
	return 0
}

func main() {
	// invocando a função soma
	// s := soma(4, 5)
	// fmt.Println(s)

	// invocando a função operacaoAritmetica
	fmt.Println(operacaoAritmetica(6, 2, Soma))
	fmt.Println(operacaoAritmetica(6, 2, Subtra))
	fmt.Println(operacaoAritmetica(6, 2, Multip))
	fmt.Println(operacaoAritmetica(6, 2, Divis))

}
