package repository

import (
	"fmt"
	"rpcprac/todo"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(name string) (todo.User, error) {

	query := `
	INSERT INTO users (name)
	VALUES ($1)
	RETURNING id, name, level, xp
	`

	var respUser todo.User

	err := r.db.QueryRow(
		query,
		name,
	).Scan(
		&respUser.Id,
		&respUser.Name,
		&respUser.Level,
		&respUser.Xp,
	)

	if err != nil {
		return todo.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return respUser, nil
}

func (r *UserPostgres) GetUser(id int64) (todo.User, error) {

	query := `
	SELECT id, name, level, xp
	FROM users
	WHERE id = $1
	`

	var respUser todo.User

	err := r.db.Get(&respUser, query, id)
	if err != nil {
		return todo.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return respUser, nil
}

func (r *UserPostgres) ListUsers() ([]todo.User, error) {

	query := `
	SELECT id, name, level, xp
	FROM users
	`

	var users []todo.User

	err := r.db.Select(&users, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return users, nil
}
