package main

/*
A mesma empresa precisa ler o arquivo armazenado, para isso exige que:
- Seja impresso na tela os valores tabelados, com título ( à esquerda para o ID e à
direita para o Preço e Quantidade), o preço, a quantidade e abaixo do preço o total
deve ser exibido (somando preço por quantidade)

Exemplo:
ID 		 Preco 			Quantidade
111223 30012.00 	1
444321 1000000.00 4
434321 50.50 			1

			 4030062.50
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	file, err := os.Open("produtos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := tabwriter.NewWriter(os.Stdout, 15, 7, 0, '\t',
		tabwriter.AlignRight)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	cabecalho := strings.Split(scanner.Text(), ",")
	for _, c := range cabecalho {
		fmt.Fprintf(w, "%s\t", c)
	}
	fmt.Fprintln(w)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for _, v := range values {
			fmt.Fprintf(w, "%s\t", v)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
}
