package repository

import (
	"context"
	"go-todo-api/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	

)

type TodoRepository struct {
    DB *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
    return &TodoRepository{DB: db}
}


func (r *TodoRepository) GetTodos(ctx context.Context) ([]model.Todo, error) {
    rows, err := r.DB.Query(ctx, "SELECT id, title, completed FROM todos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    todos := []model.Todo{}

    for rows.Next() {
        var t model.Todo
        if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
            return nil, err
        }
        todos = append(todos, t)
    }

    return todos, nil
}

func (r *TodoRepository) CreateTodo(ctx context.Context, title string, completed bool) error {
    _, err := r.DB.Exec(ctx, "INSERT INTO todos (title, completed) VALUES ($1, $2)", title, completed)
    return err
}
func (r *TodoRepository) DeleteTodo(ctx context.Context, id int) error {
    _, err := r.DB.Exec(ctx, "DELETE FROM todos WHERE id=$1", id)
    return err
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, id int, title string, completed bool) error {
    _, err := r.DB.Exec(ctx,"UPDATE todos SET title=$1, completed=$2 WHERE id=$3",title, completed, id,)
    return err
}
