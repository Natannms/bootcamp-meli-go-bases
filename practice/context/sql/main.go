package main

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var conn *sql.DB

// quando abrimos uma conexão com um banco de dados, é retornado na verdade um pool de conexões,
// que podem ser usadas entre várias funções ao mesmo tempo, desde que tenha alguma disponível
// [
//   conn,
//   conn,
//   conn,
//   conn,
//   conn,
//   conn,
//   conn,
// ]

type Produto struct {
	Id    string
	Nome  string
	Preco float64
}

func main() {
	_, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	// executando uma query no banco de dados sem contexto
	// rows, err := conn.Query("SELECT * FROM produtos")
	// executando uma query no banco de dados com contexto

}

func getProducts(ctx context.Context) *[]Produto {
	var produtos *[]Produto
	rows, err := conn.QueryContext(ctx, "SELECT id, name, preco FROM produtos")
	if err != nil {
		// fazemos um switch para tratar cada tipo de erro que uma query do banco de dados pode retornar
		switch {
		case errors.Is(err, sql.ErrNoRows):
			{
				errors.New("não encontramos valores")
			}

		default:
			errors.New("um erro desconhecido")
		}
	}

	for rows.Next() {
		var produto Produto
		rows.Scan(&produto.Id, &produto.Nome, &produto.Preco)
	}

	return produtos
}
