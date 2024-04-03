package main

import (
	"fmt"
	"os"
)

func main() {
	// print("=========================== Criar uma variavel de ambiente do sistema operacional ==========================\n")

	// criar uma variavel de ambiente do sistema
	err := os.Setenv("GO_LANG", "go")
	if err != nil {
		fmt.Println(err)
	}

	print("=========================== Retorna uma variavel de ambiente do sistema operacional ==========================\n")
	value := os.Getenv("GO_LANG")
	fmt.Printf("O valor da variavel de ambiente é: %s\n", value)

	print("=========================== RETORNA O VALOR DA VARIAVEL DE AMBIENTE E CONFIRMA SE A MESMA EXISTE ==========================\n")
	// retorna 2 valores, o nome da variavel e um boleano
	value, ok := os.LookupEnv("TESTE")
	if ok {
		fmt.Printf("O valor da variavel de ambiente é :  %s", value)
	} else {
		fmt.Println("A variavel nao existe")
	}

	print("=========================== LEITURA DE DIRETORIOS ==========================\n")
	// Retorna uma lista de diretorios
	files, err := os.ReadDir("../")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(files)

	//um loop para imprimir cada item do slice
	for _, f := range files {
		fmt.Println(f)
	}

	print("=========================== LEITURA DE ARQIVOS ==========================\n")
	// retorna o conteudo do arquivo
	data, err := os.ReadFile("./myFile.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	print("=========================== ESCRITA DE ARQUIVOS ==========================\n")
	// criamos a variavel
	d1 := []byte("Olá Mundo!")
	err = os.WriteFile("./lista.txt", d1, 0644)

	/*
	   Explicando PERMISSÕES
	   Cada grupo consiste em três bits, que podem ter valores de 0 a 7. Cada bit tem um significado específico:

	   O primeiro bit (o mais à esquerda) indica se o arquivo é um arquivo regular (0) ou um diretório (1).
	   Os três bits seguintes (do meio) representam as permissões do proprietário do arquivo.
	   Os três bits seguintes representam as permissões do grupo do arquivo.
	   Os três bits finais representam as permissões para outros usuários.
	   As permissões possíveis são:

	   0 (zero) - Sem permissão (---)
	   1 - Execução (--x)
	   2 - Escrita (-w-)
	   3 - Escrita e execução (-wx)
	   4 - Leitura (r--)
	   5 - Leitura e execução (r-x)
	   6 - Leitura e escrita (rw-)
	   7 - Leitura, escrita e execução (rwx)
	*/

}
