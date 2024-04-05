package main

import "fmt"

// essa é a estrutura da interface error
type error interface {
	Error() string
}

// exemplo de como ela é implementada
type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return "my error info"
}

func main() {
	//Errorf do pacote fmt implementa a interface nativa do go error
	//assim como o metodo MeuErrorF do exemplo
	fmt.Errorf("informação do meu erro")
}
