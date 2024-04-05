package main

import "fmt"

func main2() {
	chan1 := make(chan int)

	go func() {
		chan1 <- 1
	}()

	// estamos recendo um valor
	valorRecebido := <-chan1
	fmt.Println(valorRecebido)
}
