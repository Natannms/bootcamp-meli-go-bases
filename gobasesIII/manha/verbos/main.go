package main

import "fmt"

func main() {
	// verbos mais usados com fmt : [ %v %T %t %s %f %d %b %o %c %p ]

	// %v permite qualquer valor
	fmt.Printf("Permite qualquer %v\n", "VALOR")

	// %T exibe o tipo da variável
	var x int = 42
	fmt.Printf("Tipo da variável x: %T\n", x)

	// %t exibe o valor booleano
	isTrue := true
	fmt.Printf("O valor de isTrue é: %t\n", isTrue)

	// %s exibe uma string
	name := "John"
	fmt.Printf("Meu nome é: %s\n", name)

	// %f exibe um número de ponto flutuante
	pi := 3.14159265359
	fmt.Printf("O valor de pi é: %f\n", pi)

	// %d exibe um número inteiro
	age := 30
	fmt.Printf("Minha idade é: %d\n", age)

	// %b exibe um número em formato binário
	num := 10
	fmt.Printf("O valor de num em binário é: %b\n", num)

	// %o exibe um número em formato octal
	num2 := 8
	fmt.Printf("O valor de num2 em octal é: %o\n", num2)

	// %c exibe o caractere correspondente ao código ASCII
	char := 'A'
	fmt.Printf("O caractere correspondente ao código ASCII %d é: %c\n", char, char)

	// %p exibe um ponteiro
	ptr := &x
	fmt.Printf("O endereço de x é: %p\n", ptr)
}
