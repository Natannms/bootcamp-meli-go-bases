package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando ...")
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Fim")
}
