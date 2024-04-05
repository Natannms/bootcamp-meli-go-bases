package main

import "fmt"

/*
Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é declarar diferentes variáveis.
Ajude o professor com as seguintes questões:
Verifique quais dessas variáveis ​​declaradas pelo aluno estão corretas.
Corrigir as incorretas.

*/

func main() {
	// before
	// var 1nome string
	// var sobrenome string
	// var int idade
	// 1sobrenome := 6
	// var licenca_para_dirigir = true
	// var estatura da pessoa int
	// quantidadeDeFilhos := 2

	// after

	var nome string
	var sobrenome string
	var idade int
	sobrenome = "algum sobrenome"
	var licenca_para_dirigir = true
	var estaturaDaPessoa int
	quantidadeDeFilhos := 2

	fmt.Println(nome, sobrenome, idade, licenca_para_dirigir, estaturaDaPessoa, quantidadeDeFilhos)

}
