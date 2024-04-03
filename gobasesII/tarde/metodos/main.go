package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func (circulo Circulo) area() float64 {
	return math.Pi * circulo.raio * circulo.raio
}

func (c Circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (c *Circulo) setRaio(raio float64) {
	c.raio = raio
}

func main() {
	c := Circulo{raio: 5}
	fmt.Println(c.area())
	fmt.Println(c.perimetro())
	c.setRaio(10)
	fmt.Println(c.area())
	fmt.Println(c.perimetro())
}
