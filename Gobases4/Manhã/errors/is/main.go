package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("Erro numero 1")

func x() error {
	return fmt.Errorf("informação extra do erro: %w", err1)
	//informação extra do erro: Erro numero 1
}

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return "my error info"
}

func main() {
	e := x()
	coincidence := errors.Is(e, err1)
	fmt.Println(coincidence) //imprime true
}
