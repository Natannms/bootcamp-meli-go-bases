package main

import "fmt"

/*
A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma das
letras separadamente para soletrá-la. Para isso terão que:
Crie uma aplicação que tenha uma variável com a palavra e imprima o número de letras que ela contém.
Em seguida, imprimi cada uma das letras.

*/

func main() {
	bootcamp := "Bootcamp GO W3"

	fmt.Println("tamanho:", len(bootcamp))

	for _, r := range bootcamp {
		fmt.Println(string(r))
	}
}
