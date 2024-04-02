package ma

func main() {
	//Criando map

	// map[tipo da chave] tipo do valor
	myMap := map[string]int{

		//chave			 valor
		"primeira-chave": 25,
	}

	//adicionando uma nova chave ao map com seu valor
	myMap["segunda-chave"] = 15

	//atualizando o valor de uma chave existente do map
	myMap["primeira-chave"] = 35

	//deletando uma chave especifica do map
	delete(myMap, "primeira-chave")
}
