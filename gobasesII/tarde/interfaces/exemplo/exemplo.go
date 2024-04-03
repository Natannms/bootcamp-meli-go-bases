package main

type ICut interface {
	cut()
}

type jardineiro struct {
	Nome string
}

func (f jardineiro) cut() {}

type auxiliarGeral struct {
	Nome string
}

func (a auxiliarGeral) cut() {}

func cortarGrama(f ICut) {
	f.cut()
}

func main() {
	jardineiro := jardineiro{"Lucas"}
	auxiliarGeral := auxiliarGeral{"Vitor"}

	cortarGrama(jardineiro)
	cortarGrama(auxiliarGeral)
}
