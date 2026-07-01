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

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.db.QueryRow(ctx, "SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	err := r.db.QueryRow(ctx, "SELECT id, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
