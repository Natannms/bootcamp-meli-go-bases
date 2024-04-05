package main

import (
	"errors"
	"fmt"
)

/*
	Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar
	produtos e devolver o valor do preço total.
	As empresas têm 3 tipos de produtos:
	- Pequeno, Médio e Grande.
	Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

	Lista de custos adicionais:
	- Pequeno: O custo do produto (sem custo adicional)
	- Médio: O custo do produto + 3% pela disponibilidade no estoque
	- Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
	adicional pelo envio de $2500.

	Requisitos:
	- Criar uma estrutura “loja” que guarde uma lista de produtos.
	- Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
	- Criar uma interface “Produto” que possua o método “CalcularCusto”
	- Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
	- Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome
	e preço, e devolva um Produto.
	- Será necessário uma função “novaLoja” que retorne um Ecommerce.
	- Interface Produto:
	- Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o
	custo adicional segundo o tipo de produto.

	- Interface Ecommerce:
	- Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
	base no custo total dos produtos + o adicional citado anteriormente (caso a categoria
	tenha)
	- Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
	e adicioná-lo a lista da loja
*/

type IProduto interface {
	CalcularCusto() (float64, error)
}

type IEcommerce interface {
	Total() float64
	Adicionar(produto IProduto)
}

const (
	Pequeno = "Pequeno"
	Medio   = "Médio"
	Grande  = "Grande"
)

type Produto struct {
	Tipo  string
	Nome  string
	Preco float64
}

func (p Produto) CalcularCusto() (float64, error) {
	switch p.Tipo {
	case Pequeno:
		return p.Preco, nil
	case Medio:
		return p.Preco * 1.03, nil
	case Grande:
		return (p.Preco * 1.06) + 2500, nil
	default:
		return 0.0, errors.New("type product doesn't exist")
	}
}

type Loja struct {
	produto []IProduto
}

func (l Loja) Total() float64 {
	totalPrice := 0.0
	for _, produto := range l.produto {
		price, error := produto.CalcularCusto()
		if error != nil {
			fmt.Println(error)
		} else {
			totalPrice += price
		}
	}
	return totalPrice
}

func (l *Loja) Adicionar(produto IProduto) {
	l.produto = append(l.produto, produto)
}

func main() {
	produto_computador := novoProduto(Medio, "Computador", 10000)
	fmt.Println(produto_computador)
	loja := novaLoja(produto_computador)
	produto_celular := novoProduto(Pequeno, "Celular", 8000)
	loja.Adicionar(produto_celular)
	fmt.Println(loja)
	loja.Adicionar(novoProduto(Grande, "Televisao", 20000))
	fmt.Println(loja)
	totalPrice := loja.Total()
	fmt.Println("O preço total da loja é igual a:", totalPrice)
}

func novoProduto(typeProduct string, name string, price float64) Produto {
	return Produto{
		Nome:  name,
		Tipo:  typeProduct,
		Preco: price,
	}
}

func novaLoja(produtos ...IProduto) IEcommerce {
	return &Loja{produtos}
}
