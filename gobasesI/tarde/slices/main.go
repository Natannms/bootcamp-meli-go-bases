package main

import "fmt"

func main() {
	//slices

	//criando um slice
	var s = []string{"Natan", "Lucas", "Luan"}
	fmt.Println(s)

	//criando slice com Make
	m := make([]int, 10)
	fmt.Println(m)

	//obtendo dados do slice a partir de um range
	primes := []int{1, 2, 3, 4, 5}
	fmt.Println(primes[1:4])

	//adicionando valores ao final de um slice
	primes = append(primes, 10, 15, 16)
	fmt.Println(primes)
}
