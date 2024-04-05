package main

/*
	Uma empresa que vende produtos de limpeza precisa de:
	1. Implementar uma funcionalidade para guardar um arquivo de texto, com a informação
	de produtos comprados, separados por ponto e vírgula (csv).
	2. Deve possuir o ID do produto, preço e a quantidade.
	3. Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.
*/

import (
	"errors"
	"fmt"
	"os"
)

type produto struct {
	id         int
	quantidade int
	preco      float64
}

func (p produto) gerarLinhaCSV() string {
	return fmt.Sprintf("%d,%d,%.2f\n", p.id, p.quantidade, p.preco)
}

func (p produto) gerarCabecalhoCSV() string {
	return "id,quantidade,preco\n"
}

func gerarCsv(caminho string, produtos []produto) error {
	if len(produtos) == 0 {
		return errors.New("quantidade de produto inválida")
	}
	file, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close()
	p := produtos[0]
	if _, err = file.WriteString(p.gerarCabecalhoCSV()); err != nil {
		return fmt.Errorf("erro ao gerar cabeçalho: %w", err)
	}
	for _, p = range produtos {

		if _, err = file.WriteString(p.gerarLinhaCSV()); err != nil {
			return fmt.Errorf("erro ao salvar linha: %w", err)
		}
	}
	return nil
}

func main() {
	produtos := []produto{
		{
			id:         3,
			quantidade: 8,
			preco:      6.99,
		},
		{
			id:         2,
			quantidade: 20,
			preco:      12.99,
		},
		{
			id:         1,
			quantidade: 10,
			preco:      9.99,
		},
	}
	gerarCsv("produtos.csv", produtos)
}
