package main

import (
	"fmt"
	"log"
	"os"
)

// FIFO (First In First out) -> fila

type pessoas struct {
	Nome string
}

func readFile() ([]pessoas, error) {
	file, err := os.Open("customers.txt")
	defer file.Close()
	var pessoas []pessoas
	return pessoas, err
}

func main2() {
	// quando eu chamo uma função para ler o arquivo, o arquivo é aberto nela, e fechado assim que ela termina de executar
	// dessa forma, trabalhos no restante da função apenas com o conteúdo do que lemos do arquivo
	pessoas, err := readFile()
	if err != nil {
		log.Fatal(err)
	}
	// LIFO (Last in Fist Out)
	// aqui estamos colocando funções para serem executadas após o término do bloco em que está presente, ele funciona em forma
	// de pilha, ou seja, o último item a ser colocado, vai ser o primeiro a ser tirado para a execução do mesmo
	fmt.Println(pessoas)
	defer fmt.Println("depois 1")
	defer fmt.Println("depois 2")
}
