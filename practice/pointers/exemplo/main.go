package main

import "fmt"

/*
  | nome | endereço de memória | valor |
*/

type Pessoa struct {
	Nome string
}

func (p *Pessoa) SetNome(nome string) {
	p.Nome = nome
}

// 3 formas de se trabalhar com uma struct de pessoas
// []Pessoa
// []*Pessoa
// *[]Pessoa

func newPessoa(nome string) Pessoa {
	// criando uma variável pessoa
	pessoa := Pessoa{
		Nome: nome,
	}

	// retornando apenas o endereço de memória de pessoa
	return pessoa
}

func main() {

	// criando uma variável do tipo int
	var idade int = 45
	// aqui criamos um ponteiro que aponta para uma variável do tipo int
	var pIdade *int
	fmt.Printf("o valor do ponteiro ao declarar é %v \n", pIdade)

	// aqui estamos dando um valor para um ponteiro, que deve ser um endereço de memória
	pIdade = &idade

	fmt.Printf("o valor do ponteiro ao atribuir é %v \n", pIdade)

	// para acessar o valor da variável que o ponteiro aponta, usamos o "*"
	fmt.Printf("o valor da variável que o ponteiro aponta é %v \n", *pIdade)

	// para mudarmos o valor da variável que o ponteiro aponta, usamos o "*" para acessar o valor, e atribuimos um novo valor
	*pIdade = 50
	fmt.Printf("o novo valor da variável que o ponteiro aponta é %v \n", *pIdade)
	fmt.Printf("o valor do ponteiro ao atribuir é %v \n", pIdade)

	// criando uma struct pessoa
	p1 := newPessoa("Fulano")
	fmt.Println(p1)

	// podemos alterar o valor de qualquer atributo da struct dessa forma, ou usamos o método para isso
	*&p1.Nome = "Beltrano"
	fmt.Println(p1)

	// aqui estamos alterando o atributo via método
	p1.SetNome("Ciclano")
	fmt.Println(p1)
}
