package main

import (
	"fmt"
	"time"
)

func main() {

	println("\n========= for standard =========\n")
	// for standard

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}

	println("\n========= for range =========\n")
	// for...range
	frutas := []string{"maçã", "banana", "pêra"}

	// ignorando variável retornada do range
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// usando variáveis retornadas do range
	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}

	println("\n========= for infinito =========\n")
	// for infinito

	// sum = 0
	// for {
	// 	sum++
	// }

	sum = 0
	for {
		sum++
		if sum >= 1000 {
			break
		}
	}
	fmt.Println(sum)

	println("\n========= for while =========\n")
	// for while
	sum = 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	println("\n========= for with continue =========\n")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "is odd")
	}

	println("\n========= IF =========\n")
	// if

	// caso não definirmos os bits do int, vai ser considerado os bits do processador (32bits ou 64bits), o compilador no processo de build vai definir isso
	var number uint64 = 2
	if number%2 == 0 {
		fmt.Println("esse número é par")
	}

	println("\n========= IF...ELSE =========\n")
	// if else

	estaChovendo := true
	if estaChovendo {
		fmt.Println("hoje vamos ficar de netflix mesmo")
	} else {
		fmt.Println("hoje tem festa")
	}

	println("\n========= IF... ELSE IF ... ELSE =========\n")
	// if...if else...else

	estaChovendo = false
	filmeNovoEmCartaz := false
	if estaChovendo {
		fmt.Println("hoje vamos ficar de netflix mesmo")
	} else if filmeNovoEmCartaz {
		fmt.Println("vamos de cineminha hoje")
	} else {
		fmt.Println("hoje tem festa")
	}

	println("\n========= IF com declaração curta =========\n")
	// if com declaração curta

	if estaChovendo := true; estaChovendo {
		fmt.Println("algum bloco de código aqui...")
	}

	println("\n========= SWITCH =========\n")
	// wwitch

	status := "aprovado"

	switch status {
	case "aprovado":
		fmt.Println("parabéns!")
	default:
		fmt.Println("não conseguimos identificar seu status, tente novamente mais tarde")
	}

	// outro exemplo de switch
	dias := 1
	switch dias {
	case 0:
		fmt.Println("Segunda-feira")
	case 1:
		fmt.Println("Terça-feira")
	case 2:
		fmt.Println("Quarta-feira")
	case 3:
		fmt.Println("Quinta-feira")
	case 4:
		fmt.Println("Sexta-feira")
	case 5:
		fmt.Println("Sábado")
	case 6:
		fmt.Println("Domingo")
	default:
		fmt.Println("Desconhecido")
	}

	// switch com condicional
	var idade uint8 = 18
	switch {
	case idade >= 150:
		fmt.Println("É imortal?")
	case idade >= 18:
		fmt.Println("É maior de idade")
	default:
		fmt.Println("É menor de idade")
	}

	// switch com múltiplos cases

	dia := "domingo"
	switch dia {
	case "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira":
		fmt.Printf("%s é um dia de semana\n", dia)
	default:
		fmt.Printf("%s é um dia de final de semana\n", dia)
	}

	// switch com declaração curta
	switch dia := "sábado"; dia {
	case "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira":
		fmt.Printf("%s é um dia de semana\n", dia)
	default:
		fmt.Printf("%s é um dia do final de semana\n", dia)
	}

	// switch com fallthrough

	hoje := time.Now()
	diaDoMes := hoje.Day()
	switch diaDoMes {
	case 5, 10, 15:
		fmt.Println("Limpar a casa.")
	case 25, 26, 27:
		fmt.Println("Comprar comida.")
		fallthrough
	case 31:
		fmt.Println("Hoje tem festa.")
	default:
		fmt.Println("Não há informações disponíveis para esse dia.")
	}

	// podemos usar o break dentro do bloco do switch se ele estiver dentro de um for
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Printf("%d é par\n", i)
			break
		case i == 7 || i == 3:
			fmt.Printf("%d é primo\n", i)
			fallthrough
		case i%2 == 1:
			fmt.Printf("%d é impar\n", i)
			break
		default:
			fmt.Println("é um número mágico")
		}
	}
}
