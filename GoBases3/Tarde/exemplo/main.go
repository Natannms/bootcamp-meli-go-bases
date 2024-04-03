package main

import "fmt"

func main() {
	var v int = 19
	var p *int
	// Criamos um ponteiro p, para referenciar o endereço na memória
	// variável v.
	p = &v
	fmt.Printf("O ponteiro p refere-se ao endereço de memória: %v \n", p)
	fmt.Printf("Ao desreferenciar o ponteiro p, obtenho o valor: %d \n", *p)
}
