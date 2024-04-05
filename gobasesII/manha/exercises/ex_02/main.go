package main

import "fmt"

/*
Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
qual possamos passar N quantidade de números inteiros e devolva a média.
Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
*/

func calculaMedia(notas ...float64) float64 {
	sum := 0.0
	for _, n := range notas {
		sum += n
	}
	return sum / float64(len(notas))
}

func main() {
	fmt.Printf("Calcula Média: %.2f\n", calculaMedia(10, 8, 9))
}
