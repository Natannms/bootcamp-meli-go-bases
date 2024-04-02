package main

import (
	"fmt"

	"github.com/Natanms/meu-projeto/server"
)

func main() {
	// Decalarações de variaveis
	// var name string
	// var(
	// 	idade = 20
	// 	cidade = "Belo Horizonte"
	// )
	//  const palavraSecreta =  "SECRET"

	//Tipos
	// var status int
	// var peso float64
	// var texto string
	// var boleano bool
	// var byte byte

	//scapes
	// fmt.Println("Ola \n Mundo")
	// fmt.Println("Ola \\ Mundo")
	// fmt.Println(`Ola ' " Mundo`)

	//importações e exportações
	key := server.KEY

	//não é possivel importar a varaivel profile

	perfil := server.profile //server.profile é undefined pois não é exportavel

	fmt.Println(key)
	fmt.Println(perfil)
}
