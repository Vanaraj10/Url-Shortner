package repository

import (
	"context"

	"github.com/Vanaraj10/Url-Shortner/models"
)

type URLRepository interface {
	Create(ctx context.Context, url *models.URL) error
	FindByShort(ctx context.Context, short string) (*models.URL, error)
	FindByUser(ctx context.Context, userID int64) ([]*models.URL, error)
}
