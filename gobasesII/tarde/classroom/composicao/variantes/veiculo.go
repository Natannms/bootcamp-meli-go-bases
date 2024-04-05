package variantes

import "fmt"

type Veiculo struct {
	km   float64
	hora float64
}

func (v Veiculo) detalhe() {
	fmt.Printf("km:\t%.2f\nhora:\t%.2f\n", v.km, v.hora)
}
