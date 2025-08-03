package repository

import (
	"context"

	"github.com/Vanaraj10/Url-Shortner/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userPostgres struct {
	db *pgxpool.Pool
}

func NewUserPostgres(db *pgxpool.Pool) *userPostgres {
	return &userPostgres{db: db}
}

func (r *userPostgres) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, created_at`
	return r.db.QueryRow(ctx, query, user.Username, user.Password).Scan(&user.ID, &user.CreatedAt)
}

func (r *userPostgres) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT id, username, password, created_at FROM users WHERE username = $1`
	row := r.db.QueryRow(ctx, query, username)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userPostgres) FindByID(ctx context.Context, id int64) (*models.User, error) {
	query := `SELECT id, username, password, created_at FROM users WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)
	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}
