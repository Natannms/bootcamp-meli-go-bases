package main

/*
	1. Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do
	tipo “int”.
	2. Crie um erro personalizado com uma struct que implemente “Error()” com a
	mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
	disparado quando “salário” for menor que 15.000. Caso contrário, imprima no
	console a mensagem “Necessário pagamento de imposto”.
*/

type MyError struct{}

func (m *MyError) Error() string {
	return "erro: O salário digitado não está dentro do valor mínimo"
}

func main() {
	salario := 5_000

	if salario < 15000 {
		err := &MyError{}
		println(err.Error())
	} else {
		println("Necessário pagamento de imposto")
	}
}
