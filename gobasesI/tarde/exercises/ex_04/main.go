package main

import "fmt"

/*
Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários.
De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

Por outro lado, você também precisa:
Saiba quantos de seus funcionários têm mais de 21 anos.
Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
Excluir Pedro do mapa.

*/

func main() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	// idade de Benjamin
	fmt.Println("idade de Benjamin:", employees["Benjamin"])

	// funcionários com mais de 21
	c := 0
	for _, employee := range employees {
		if employee > 21 {
			c++
		}
	}
	fmt.Printf("funcionários com mais de 21 anos: %d\n", c)

	// adicionar Frederico
	fmt.Println("##### adicionar Frederico")
	fmt.Println("antes:", employees)
	employees["Frederico"] = 25
	fmt.Println("depois:", employees)

	// excluir Pedro
	fmt.Println("##### excluir Pedro")
	fmt.Println("antes:", employees)
	delete(employees, "Pedro")
	fmt.Println("depois:", employees)

}
