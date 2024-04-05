package variantes

import "fmt"

// Automóvel é um veículo
type Automovel struct {
	v Veiculo
}

func (a *Automovel) Correr(minutos int) {
	a.v.hora = float64(minutos) / 60
	a.v.km = a.v.hora * 100
}

func (a Automovel) Detalhe() {
	fmt.Println("\nV:\tAutomovel")
	a.v.detalhe()
}
