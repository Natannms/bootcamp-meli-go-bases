package main

/*
	Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
	produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o
	mesmo endereço de memória no main do programa e nas funções.

	Estruturas necessárias:
	- Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
	- Produto: Nome, preço, quantidade.
	Algumas funções necessárias:
	- Novo produto: recebe nome e preço, e retorna um produto.
	- Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona
	o produto ao usuário.
	- Deletar produtos: recebe um usuário, apaga os produtos do usuário.
*/

import "fmt"

type User struct {
	Name     string
	LastName string
	Email    string
	Products []*Product // array de produtos
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}

func AddProduct(user *User, product *Product, quantity int) {
	// Verifica se o produto já existe na lista do usuário
	for _, p := range user.Products {
		if p == product {
			p.Quantity += quantity
			return
		}
	}

	// Se o produto não existe, adiciona um novo produto à lista do usuário
	product.Quantity = quantity
	user.Products = append(user.Products, product)
}

func DeleteProducts(user *User) {
	user.Products = nil
}

func main() {
	// Cria um novo usuário
	user := &User{
		Name:     "João",
		LastName: "Silva",
		Email:    "joao.silva@example.com",
	}

	// Adiciona um produto ao usuário
	product1 := NewProduct("Celular", 1000.0)
	AddProduct(user, product1, 1)

	// Adiciona outro produto ao usuário
	product2 := NewProduct("Notebook", 2500.0)
	AddProduct(user, product2, 2)

	// Mostra a lista de produtos do usuário
	fmt.Printf("Produtos de %s %s:\n", user.Name, user.LastName)
	for _, p := range user.Products {
		fmt.Printf("- %s (R$ %.2f) x %d\n", p.Name, p.Price, p.Quantity)
	}

	// Apaga todos os produtos do usuário
	DeleteProducts(user)

	// Mostra a lista de produtos do usuário (deve ser vazia)
	fmt.Printf("Produtos de %s %s:\n", user.Name, user.LastName)
	for _, p := range user.Products {
		fmt.Printf("- %s (R$ %.2f) x %d\n", p.Name, p.Price, p.Quantity)
	}
}
