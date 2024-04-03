package main

import "fmt"

type Pessoa struct {
	Nome    string
	Salario float64
}

func (p *Pessoa) setSalario(novoSalario float64) {
	p.Salario = novoSalario
	fmt.Printf("Novo salário do lucas: %.2f \n", novoSalario)
}

func (p Pessoa) Print() {
	fmt.Printf("Nome: %s \nSalário: %.2f\n", p.Nome, p.Salario)
}

func NewPessoa(name string, salario float64) Pessoa {
	return Pessoa{
		Nome:    name,
		Salario: salario,
	}
}

func main() {
	p1 := Pessoa{"Lucas", 0}
	fmt.Printf("Salário %.2f do %s \n", p1.Salario, p1.Nome)

	p1.setSalario(100_000)

	fmt.Printf("Novo salário %.2f do %s \n", p1.Salario, p1.Nome)

	p1.Print()
}
