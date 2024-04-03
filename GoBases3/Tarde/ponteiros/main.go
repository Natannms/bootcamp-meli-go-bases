package main

import "fmt"

func main() {
	var p *int

	//criando ponteiros com new()
	var p2 = new(int)

	//atribuindo a uma variavel um ponteiro usando
	p3 := &p2

	fmt.Println(p, p2, p3)
}
