package main

import "fmt"

type Pessoa struct {
	Nome      string
	Genero    string
	Idade     int
	Profissao string
	Peso      float64
	Gostos    Preferencias
}

type Preferencias struct {
	Comidas  string
	Filme    string
	Series   []string
	Animes   string
	Esportes string
}

func main() {
	p1 := Pessoa{"Paula", "F", 34, "Engenheira", 65.5, Preferencias{}}

	fmt.Println(p1)

	p2 := Pessoa{
		Nome:      "Vitor",
		Idade:     30,
		Profissao: "Arquiteto",
		Peso:      77,
	}

	fmt.Println(p2)

	// acessando propriedades da struct

	fmt.Println("\npessoa 1\n")
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	fmt.Println("\npessoa 2\n")
	fmt.Println(p2.Nome)
	fmt.Println(p2.Idade)

	p2.Idade = 31
	fmt.Println(p2.Idade)

	var p3 Pessoa
	p3.Nome = "Fernando"
	p3.Idade = 15

	fmt.Println("\npessoa 3\n")
	fmt.Println(p3.Nome)
	fmt.Println(p3.Idade)

	fmt.Printf("printando chave e valor %+v \n", p3)

	erika := Pessoa{
		Nome:      "Erika",
		Idade:     23,
		Genero:    "F",
		Profissao: "Software Developer",
		Gostos: Preferencias{
			Filme:   "Scarface",
			Comidas: "Tapioca",
			Series: []string{
				"Mr. Robot",
			},
			Animes:   "South Park",
			Esportes: "Ciclismo",
		},
	}

	fmt.Println(erika)

	lucas := Pessoa{
		Nome:      "Lucas",
		Idade:     21,
		Genero:    "M",
		Profissao: "Engenheiro de software",
		Peso:      63,
		Gostos: Preferencias{
			Comidas:  "File de frango a milenaesa com batata frita",
			Filme:    "WALL-E",
			Series:   []string{"The big bang theory"},
			Animes:   "Charlotte",
			Esportes: "Futebol",
		},
	}

	fmt.Println("Mudando os gostos da p3")
	lucas.Gostos.Filme = "9 a salvação"
	fmt.Println(lucas)

}
