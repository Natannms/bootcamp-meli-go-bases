package main

import "fmt"

func main() {

	println("====== Utilizando o print nativo do go ======\n")
	// Usamos quando vamos printar algum tipo nativo do golang

	print("Hello \n")
	println("Hello")
	println(1)
	println(true)
	println([]int{1, 2, 3})

	// Produz um erro ao executar, o go não consegue entender a complexidade de uma struct
	// println(struct{ Nome string }{"Lucas"})

	println("\n====== Utilizando o pacote fmt do go ======\n")
	// Quando vamos usar tipos complexos do golang ou precisamos fazer alguma formatação

	fmt.Println([]int{1, 2, 3})
	fmt.Println(struct{ Nome string }{"Lucas"})

	// Métodos úteis do pacote fmt

	// printar no console um valor sem quebra de linha
	fmt.Print("Hello \n")
	fmt.Println("Hello")

	// Usando o método print para retornar uma string
	println(fmt.Sprint("Retornando Hello"))

	// Usando o método print para retornar uma string formatada
	println(fmt.Sprintf("Hello %s", "Alunos"))

	fmt.Printf("erro: %v", fmt.Errorf("não foi possível encontrar um desenvolvedor bom usando %s", "windows phone"))
}
