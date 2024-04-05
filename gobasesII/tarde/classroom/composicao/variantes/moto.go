package variantes

import "fmt"

type Moto struct {
	v Veiculo
}

func (m *Moto) Correr(minutos int) {
	m.v.hora = float64(minutos) / 60
	m.v.km = m.v.hora * 80
}

func (m Moto) Detalhe() {
	fmt.Println("\nV:\tMoto")
	m.v.detalhe()
}
