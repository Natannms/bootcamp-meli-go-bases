package main

import "fmt"

func Aumentar(v *int) {
	*v++
}

func main() {
	var v int = 19

	Aumentar(&v)

	fmt.Println("O valor de v agora é:", v)
}
