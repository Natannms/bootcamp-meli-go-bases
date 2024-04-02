package main

import (
	"fmt"
	"importacoes_exportacoes/meupacote"
)

func main() {
	//importações e exportações
	key := meupacote.KEY // meupacote.KEY pode ser utilizado pois é uma CONSTANTE log é exportavel
	//url := meupacote.url     // meupacote.url é indefinido pois a variavel é privada, com inicial minúscula
	owner := meupacote.Owner //  meupacote.owner pode ser utilizado pois étem sua inicial maiúscula e é exportavel

	fmt.Println(owner)
	//fmt.Println(url) //# Erro: pois url é indefinido
	fmt.Println(key)
}
