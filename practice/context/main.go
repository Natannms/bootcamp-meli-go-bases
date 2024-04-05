package main

import "context"

/*
contexto são caixas que podem guardar informações ou outras caixas(contextos)
*/

func main2() {
	// criamos um contexto base vazio
	ctx := context.Background()
	// criamos um novo contexto do tipo cancel passando o contexto base
	_, cancel := context.WithCancel(ctx)
	// temos acesso a uma função chamada cancel que vai publicar uma mensagem no canal que é retornado pelo ctx.Done()
	cancel()

	<-ctx.Done()
}
