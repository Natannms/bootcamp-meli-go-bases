package main

import (
	"errors"
	"fmt"
)

/*
Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
de suas notas.
Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
indicado na função anterior
*/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calcMin(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min, nil
}
func calcMax(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max, nil
}
func calcAvg(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("input is required")

	}
	avg := 0.0
	for _, v := range values {
		avg += v
	}
	return avg / float64(len(values)), nil
}
func operation(t string) (func(values ...float64) (float64, error),
	error) {
	if t == minimum {
		return calcMin, nil
	}
	if t == average {
		return calcAvg, nil
	}
	if t == maximum {
		return calcMax, nil
	}
	return nil, errors.New("invalid function type")
}

func main() {

	minhaFunc, _ := operation(minimum)
	averageFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	min, _ := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("min: %.2f\n", min)
	avg, _ := averageFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("avg: %.2f\n", avg)
	max, _ := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("max: %.2f\n", max)
}
