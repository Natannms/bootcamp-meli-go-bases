package main

import "fmt"

/*
Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido.
Apenas concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano de atividade.
Dentro dos empréstimos que concede, não cobra juros para quem tem um salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis ​​e que imprima uma mensagem de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.
*/

func main() {
	clientes := []map[string]interface{}{
		{
			"Nome":      "João",
			"Idade":     21,
			"Atividade": 1,
			"Salário":   100_000,
		},
		{
			"Nome":      "José",
			"Idade":     25,
			"Atividade": 1,
			"Salário":   100_000,
		},
		{
			"Nome":      "Carlos",
			"Idade":     27,
			"Atividade": 2,
			"Salário":   99_000,
		},
		{
			"Nome":      "Antônio",
			"Idade":     35,
			"Atividade": 5,
			"Salário":   150_000,
		},
	}

	for _, client := range clientes {
		nome := client["Nome"]
		fmt.Println("Cliente:", nome)

		idade := client["Idade"].(int) //cast para int
		if idade <= 22 {
			fmt.Println("Não possui empréstimo disponível (idade insuficiente)")
			continue
		}

		atividade := client["Atividade"].(int) //cast para int
		if atividade <= 1 {
			fmt.Println("Não possui empréstimo disponível (atividade insuficiente)")
			continue
		}

		salario := client["Salário"].(int) //cast para int
		if salario > 100000 {
			fmt.Println("Possui empréstimo disponível sem juros")
		} else {
			fmt.Println("Possui empréstimo disponível com juros")
		}
	}

}
