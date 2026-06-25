package repository

import (
	"context"


	"github.com/jackc/pgx/v5"
)

func CreateTodo(ctx context.Context,conn *pgx.Conn,title string,completed bool,) error {

	_,err := conn.Exec(ctx,"INSERT INTO todos (title, completed) VALUES ($1, $2)",title,completed)

	return err
}