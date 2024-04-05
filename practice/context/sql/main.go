package main

import (
	"context"
	"fmt"
	"log"
	"sql/internal/produtos"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	// executando uma query no banco de dados sem contexto
	// rows, err := conn.Query("SELECT * FROM produtos")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// aqui eu poderia utilizar os valores das linhas
	//... algum uso de linha ...
	// fmt.Println(rows)

	// executando uma query no banco de dados com contexto
	produtos, err := produtos.GetProducts(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if produtos == nil {
		fmt.Println("não há produtos")
	}

}
