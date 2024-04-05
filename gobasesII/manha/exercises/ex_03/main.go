package main

import (
	"errors"
	"fmt"
)

/*
	Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
	de horas trabalhadas por mês e na categoria do funcionário.
	Se a categoria for C, seu salário é de R$1.000 por hora
	Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h
	mensais
	Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h
	mensais

	Calcule o salário dos funcionários conforme as informações abaixo:
	- Funcionário de categoria C: 162h mensais
	- Funcionário de categoria B: 176h mensais
	- Funcionário de categoria A: 172h mensais
*/

func calculaSalario(categoria string, horas int) (float64, error) {
	switch categoria {
	case "C":
		return float64(horas) * 1000, nil
	case "B":
		salario := float64(horas) * 1500
		if horas > 160 {
			return salario * 1.2, nil
		}
		return salario, nil
	case "A":
		salario := float64(horas) * 3000
		if horas > 160 {
			return salario * 1.5, nil
		}
		return salario, nil
	default:
		return 0.0, errors.New("Categoria Inválida")
	}
}

func main() {
	salario, _ := calculaSalario("C", 162)
	fmt.Printf("Salario 01: %.2f\n", salario)
	salario, _ = calculaSalario("B", 176)
	fmt.Printf("Salario 02: %.2f\n", salario)
	salario, _ = calculaSalario("A", 172)
	fmt.Printf("Salario 03: %.2f\n", salario)
}
