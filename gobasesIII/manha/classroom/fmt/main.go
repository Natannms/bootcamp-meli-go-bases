package main

import "fmt"

func main() {
	name, age := "Lucas", 22

	print("=========================== COM VARIOS PARAMETROS ==========================\n")
	// recebe n parametros
	fmt.Print(name, " tem", age, " anos de idade \n")

	print("=========================== COM QUEBRA DE LINHA ==========================\n")
	//Println
	fmt.Println(name, " tem", age, " anos de idade")

	print("=========================== COM QUANTIDADE APÓS A VÍRGULA ==========================\n")
	// print com formatação
	fmt.Printf("com limite após a virgula, %10.2f \n", 12222.123123) // exibe com limite de quantidade de numeros após a virgula
	fmt.Printf("com limite após a virgula, %.2f \n", 12222.123123)   // exibe com limite de quantidade de numeros após a virgula

	print("=========================== COM VERBOS DE IMPRESSÃO ==========================\n")
	//print com formatação usando verbos de impressão
	fmt.Printf("%s tem %d anos de idade.\n", name, age)

	print("=========================== COM QUANTIDADE DE CARACTERES ==========================\n")
	// print com definição de quantidade de caracteres
	fmt.Printf("%10d \n", 12222)
	fmt.Printf("%10s \n", "aa")

	print("=========================== CONCATENANDO VARIAVEIS COM SPRINT ==========================\n")
	//concatenando todas as variaveis usando Sprint
	concat := fmt.Sprint(name, " tem ", age, " anos de idade \n")
	fmt.Print(concat)

	print("=========================== CONCATENANDO VARIAVEIS COM SPRINT E VERBOS DE IMPRESSÃO ==========================\n")
	//usando verbos para formatar a string
	concatVerbo := fmt.Sprintf("%s tem %d anos de idade", name, age)
	fmt.Print(concatVerbo)

}
