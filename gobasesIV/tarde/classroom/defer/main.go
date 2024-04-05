package main

import (
	"fmt"
)

func teste() {

}
func main() {
	//aplicamos defer
	defer func() {
		fmt.Println("Esta função será executada mesmo gerando um panic")
	}()

	//criamos um panic
	panic("Panic")
	// log.Fatal("log fatal")
}
