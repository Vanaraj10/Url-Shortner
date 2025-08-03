package repository

import (
	"context"

	"github.com/Vanaraj10/Url-Shortner/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type urlPostgres struct {
	db *pgxpool.Pool
}

func NewURLPostgres(db *pgxpool.Pool) *urlPostgres {
	return &urlPostgres{db: db}
}

func (r *urlPostgres) Create(ctx context.Context, url *models.URL) error {
	query := `INSERT INTO urls (user_id, original, short) VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.db.QueryRow(ctx, query, url.UserID, url.Original, url.Short).Scan(&url.ID, &url.CreatedAt)
}

func (r *urlPostgres) FindByShort(ctx context.Context, short string) (*models.URL, error) {
	query := `SELECT id, user_id, original, short, created_at FROM urls WHERE short = $1`
	row := r.db.QueryRow(ctx, query, short)

	var url models.URL

	err := row.Scan(&url.ID, &url.UserID, &url.Original, &url.Short, &url.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *urlPostgres) FindByUser(ctx context.Context, userID int64) ([]*models.URL, error) {
	query := `SELECT id, user_id, original, short, created_at FROM urls WHERE user_id = $1`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []*models.URL
	for rows.Next() {
		var url models.URL
		if err := rows.Scan(&url.ID, &url.UserID, &url.Original, &url.Short, &url.CreatedAt); err != nil {
			return nil, err
		}
		urls = append(urls, &url)
	}
	return urls, nil
}
