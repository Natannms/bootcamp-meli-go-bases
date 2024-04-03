package main

import v "composicao/variantes"

func main() {
	automovel := v.Automovel{}
	automovel.Correr(360)
	automovel.Detalhe()

	moto := v.Moto{}
	moto.Correr(360)
	moto.Detalhe()
}
