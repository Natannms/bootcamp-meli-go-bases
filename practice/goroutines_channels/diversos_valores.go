package main

import (
	"fmt"
	"time"
)

func main3() {
	chan1 := make(chan int)
	defer close(chan1)

	go func() {
		i := 0
		for {
			chan1 <- i
			i++
			time.Sleep(time.Second)
		}
	}()

	go func() {
		i := 1000
		for {
			chan1 <- i
			i--
			// quando temos alguma execução na go routine que demore, por exemplo,
			// uma chamada de I/O ou uma chamada de rede, o scheduler do go passa a prioridade de execução para outra go routine
			time.Sleep(time.Second)
		}
	}()

	// valor <- chan1 <- valor

	// estamos recendo vários valores
	for msg := range chan1 {
		fmt.Println(msg)
	}
}
