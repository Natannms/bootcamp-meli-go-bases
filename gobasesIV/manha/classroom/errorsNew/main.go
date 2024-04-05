package main

import (
	"errors"
	"fmt"
)

func main() {
	statusCode := 404
	if statusCode > 399 {
		fmt.Println(errors.New("A requisição falhou !"))
	}

	fmt.Println("O programa foi encerrado corretamente!")
}
