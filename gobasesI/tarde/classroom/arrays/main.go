package main

import "fmt"

func main() {
	//Arrays

	//declaração de um array.
	//[tamanho do array] tipo dos valores do array
	var alunos [2]string

	fmt.Println(alunos)

	alunos[0] = "Luan"
	alunos[1] = "Natan"

	fmt.Println(alunos)
	fmt.Println(alunos[0])
}
