package main

import "fmt"

func main() {
	var name string
	var idade int

	name = "joão"
	idade = 32

	fmt.Println(&name)
	fmt.Println(&idade)

	x := 32
	y := x
	c := &x

	fmt.Println("endereço de x : ", &x)
	fmt.Println("endereço de y : ", &y)
	x = 10

	fmt.Println("\n\n VALOR de x : ", x)
	fmt.Println("VALOR de y : ", y)
	fmt.Println("VALOR de C : ", *c)
	*c = 20

	fmt.Println("VALOR de C : ", *c)
	fmt.Println("VALOR de x : ", x)
	fmt.Println("VALOR de y : ", y)

	fmt.Println("VALOR de y : ", &y)
	fmt.Println("VALOR de y : ", *c)
}
