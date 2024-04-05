package main

import (
	"errors"
	"fmt"
)

const (
	Soma   = "+"
	Subtra = "-"
	Multip = "*"
	Divis  = "/"
)

func minhaFuncaoComMultiRetorno(valor1, valor2 float64) (float64, string, int, bool) {
	fmt.Println(valor1, valor2)
	return 0.0, "", 0, true
}

func minhaFuncao(valores ...float64) float64 {
	resultado := 0.0
	for _, value := range valores {
		resultado += value
	}
	return resultado
}

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

func retornoOperacao(
	valores []float64,
	operacao func(value1, value2 float64) float64,
) float64 {
	resultado := 0.0
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacao(resultado, valor)
		}
	}

	return resultado
}

func operacaoAritmetica(
	operador string,
	valores ...float64,
) float64 {
	switch operador {
	case Soma:
		return retornoOperacao(valores, opSoma)
	case Subtra:
		return retornoOperacao(valores, opSubtra)
	case Multip:
		return retornoOperacao(valores, opMultip)
	case Divis:
		return retornoOperacao(valores, opDivis)
	}

	return 0
}

func operacoes(
	valor1,
	valor2 float64,
) (
	soma float64,
	subtra float64,
	multip float64,
	divis float64,
) {
	soma = valor1 + valor2
	subtra = valor1 - valor2
	multip = valor1 * valor2

	if valor2 != 0 {
		divis = valor1 / valor2
	}

	return
}

func divisao(dividendo, divisor float64) (float64, error) {

	if divisor == 0 {
		return 0, errors.New("O divisor não pode ser zero")
	}

	return dividendo / divisor, nil
}

func main() {
	// Somando valores com ellipis
	// resultado := minhaFuncao(2, 3, 2, 1, 2, 3, 4, 5, 6)
	// fmt.Printf("resultado: %.2f \n", resultado)

	// operação aritmética
	println("\n========= operação aritmética =========\n")
	resultado := operacaoAritmetica(Soma, 2, 3, 2, 1, 2, 3, 4, 5, 6)
	fmt.Printf("resultado: %.2f \n", resultado)

	println("\n========= função com multi retorno =========\n")
	// função com multi retorno
	s, r, m, d := operacoes(6, 2)

	fmt.Println("Soma:\t\t", s)
	fmt.Println("Subtracao:\t\t", r)
	fmt.Println("Multiplicacao:\t", m)
	fmt.Println("Divisao:\t", d)

	println("\n========= função com retorno de erro =========\n")
	res, err := divisao(2, 0)

	if err != nil {
		fmt.Println("algum erro aconteceu")
	} else {
		fmt.Println("resultado", res)
	}

}
