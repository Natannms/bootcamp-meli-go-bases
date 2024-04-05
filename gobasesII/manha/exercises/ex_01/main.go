package main

import "fmt"

/*
Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de
depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o
imposto de um salário.
Temos a informação que um dos funcionários ganha um salário de R$50.000 e será
descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse
caso será descontado mais 10%.
*/

func calculaImposto(salario float64) float64 {
	if salario <= 50000.00 {
		return 0.0
	} else if salario <= 150000.00 {
		return salario * 0.17
	}
	return salario * 0.27
}

func main() {
	fmt.Printf("imposto de até R$50.000: %.2f\n",
		calculaImposto(50000))
	fmt.Printf("imposto de até R$150.000: %.2f\n",
		calculaImposto(150000))
	fmt.Printf("imposto mais que R$150.000: %.2f\n",
		calculaImposto(150001))
}
