package main

import "fmt"

/*
	Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.
	Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total
	dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade,
	eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

	Precisamos de 3 estruturas:
	- Produtos: nome, preço, quantidade.
	- Serviços: nome, preço, minutos trabalhados.
	- Manutenção: nome, preço.
	Precisamos de 3 funções:
	- Somar Produtos: recebe um array de produto e devolve o preço total (preço *
	quantidade).
	- Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média
	hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se
	tivesse trabalhado meia hora).
	- Somar Manutenção: recebe um array de manutenção e devolve o preço total.

	Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na
	tela (somando o total dos 3).
*/

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Service struct {
	Name              string
	Price             float64
	MinutesWorked     int
	AverageHourlyRate float64
}

type Maintenance struct {
	Name  string
	Price float64
}

func SumProducts(products []Product, ch chan float64) {
	total := 0.0
	for _, p := range products {
		total += p.Price * float64(p.Quantity)
	}
	ch <- total
}

func SumServices(services []Service, ch chan float64) {
	total := 0.0
	for _, s := range services {
		minutes := s.MinutesWorked
		if minutes < 30 {
			minutes = 30
		}
		total += s.Price * (float64(minutes) / 60) * s.AverageHourlyRate
	}
	ch <- total
}

func SumMaintenance(maintenance []Maintenance, ch chan float64) {
	total := 0.0
	for _, m := range maintenance {
		total += m.Price
	}
	ch <- total
}

func main() {
	// criando dados de exemplo
	products := []Product{
		{Name: "Produto 1", Price: 10.0, Quantity: 3},
		{Name: "Produto 2", Price: 15.0, Quantity: 2},
	}
	services := []Service{
		{Name: "Serviço 1", Price: 50.0, MinutesWorked: 45, AverageHourlyRate: 100.0},
		{Name: "Serviço 2", Price: 100.0, MinutesWorked: 20, AverageHourlyRate: 80.0},
	}
	maintenance := []Maintenance{
		{Name: "Manutenção 1", Price: 200.0},
		{Name: "Manutenção 2", Price: 150.0},
	}

	chProducts := make(chan float64)
	chServices := make(chan float64)
	chMaintenance := make(chan float64)

	// iniciando go routines
	go SumProducts(products, chProducts)
	go SumServices(services, chServices)
	go SumMaintenance(maintenance, chMaintenance)

	total := 0.0
	total += <-chProducts
	total += <-chServices
	total += <-chMaintenance

	// mostrando resultado final
	fmt.Printf("Preço total: %.2f\n", total)
}
