package main

import "fmt"

func main() {
	animals := []string{
		"Vaca",     //0
		"Cachorro", //1
		"Gato",     //2
	}

	fmt.Println("O animal é:", animals[len(animals)])
}
