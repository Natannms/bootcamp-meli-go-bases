package produtos

import (
	"context"
	"database/sql"
	"errors"
	db "sql/database"
	"sql/internal/entities"
)

func GetProducts(ctx context.Context) (*[]entities.Produto, error) {
	var produtos *[]entities.Produto
	rows, err := db.Conn.QueryContext(ctx, "SELECT id, name, preco FROM produtos")
	if err != nil {
		// fazemos um switch para tratar cada tipo de erro que uma query do banco de dados pode retornar
		switch {
		case errors.Is(err, sql.ErrNoRows):
			e := errors.New("não encontramos valores")
			return nil, e
		default:
			errors.New("um erro desconhecido")
		}
	}

	for rows.Next() {
		var produto entities.Produto
		rows.Scan(&produto.Id, &produto.Nome, &produto.Preco)
	}

	return produtos, nil
}
