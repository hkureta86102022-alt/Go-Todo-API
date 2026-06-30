package repository

import (
	"context"
	"go-todo-api/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(DB *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: DB}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (email,password) VALUES ($1, $2)", user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
