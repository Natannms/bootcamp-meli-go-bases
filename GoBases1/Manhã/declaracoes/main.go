package main

import "fmt"

func main() {
	// Decalarações de variaveis
	var name string
	var (
		idade  = 20
		cidade = "Belo Horizonte"
	)
	const palavraSecreta = "SECRET"

	fmt.Println(name, idade, cidade, palavraSecreta)
}
