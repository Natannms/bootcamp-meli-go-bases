package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int)
	defer close(chan1)

	go func() {
		i := 0
		for {
			fmt.Println("enviando valor", i)
			i++
			chan1 <- i
			time.Sleep(time.Second)

		}
	}()

	// estou rebendo um único valor
	valor := <-chan1
	fmt.Println("valor recebido", valor)
	time.Sleep(time.Second * 5)
	fmt.Println("o programa terminou")
}
