package main

import (
	"fmt"
	"time"
)

// fluxo do canal

// valor <- canal <- valor

// exemplo de load balancer
func main() {
	messageChannel := make(chan int)

	// aqui estou criando 10 instâncias para executar alguma lógica
	for i := 0; i < 10; i++ {
		go instancia(i, messageChannel)
	}

	// enviando 100K de mensagens para o canal
	for i := 0; i < 100_000; i++ {
		messageChannel <- i
	}

	// enviando mensagem para o canal
	messageChannel <- 1_000
	// recebendo mensagem do canal
	variable := <-messageChannel
	fmt.Println(variable)
}

// o chan é uma FIFO, quem chegar primeiro vai sendo processado primeiro

// essa função vai funcionar numa go routine que vai concorrer para processar as mensagens que a main envia para o canal
func instancia(id int, messageChannel chan int) {
	for msg := range messageChannel {
		println("ID:", id, "Message:", msg)
		time.Sleep(time.Second)
	}
}
