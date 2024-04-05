package main

import (
	"fmt"
	"time"
)

/*
	Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os
	detalhes dos dados de cada um deles, conforme o exemplo abaixo:
	Nome: [Nome do aluno]
	Sobrenome: [Sobrenome do aluno]
	RG: [RG do aluno]
	Data de admissão: [Data de admissão do aluno]

	Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos
	alunos.
	Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG,
	Data e que tenha um método de detalhamento
*/

type Aluno struct {
	Nome      string
	Sobrenome string
	RG        string
	Data      time.Time
}

func (a Aluno) detalhamento() {
	fmt.Printf(`
    Nome: [%s]
    Sobrenome: [%s]
    RG: [%s]
    Data de admissão: [%s]
  `, a.Nome, a.Sobrenome, a.RG, a.Data)
}

type Universidade struct {
	alunos []Aluno
}

func (u *Universidade) AdicionarAluno(aluno Aluno) {
	u.alunos = append(u.alunos, aluno)
}

func main() {
	u1 := Universidade{}

	u1.AdicionarAluno(Aluno{
		Nome:      "Lucas",
		Sobrenome: "Alves",
		RG:        "123156",
		Data:      time.Date(2002, time.August, 22, 0, 0, 0, 0, &time.Location{}),
	})

	for _, aluno := range u1.alunos {
		aluno.detalhamento()
	}
}
