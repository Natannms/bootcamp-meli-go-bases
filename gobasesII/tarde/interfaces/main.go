package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

func detalhe(c geometria) {
	fmt.Printf("%+v \n", c)
	fmt.Println(c.area())
	fmt.Println(c.perimetro())
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func novoCirculo(raio float64) circulo {
	return circulo{raio: raio}
}

type triangulo struct {
	largura, altura float64
}

func (t triangulo) area() float64 {
	return t.largura * t.altura
}

func (t triangulo) perimetro() float64 {
	return 2*t.largura + 2*t.altura
}

func novoTriangulo(altura, largura float64) triangulo {
	return triangulo{
		altura,
		largura,
	}
}

const (
	tipoTriangulo = "triangulo"
	tipoCirculo   = "circulo"
)

func novaFigura(geoTipo string, values ...float64) geometria {
	switch geoTipo {
	case tipoTriangulo:
		return novoTriangulo(values[0], values[1])
	case tipoCirculo:
		return novoCirculo(values[0])
	}
	return nil
}

func main() {
	// utilizando interfaces

	// c1 := novoCirculo(10)
	// t1 := novoTriangulo(10, 10)

	c1 := novaFigura("circulo", 10)
	t1 := novaFigura("triangulo", 10)

	detalhe(c1)
	detalhe(t1)
}
