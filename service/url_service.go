package service

import (
	"context"
	"errors"

	"github.com/Vanaraj10/Url-Shortner/models"
	"github.com/Vanaraj10/Url-Shortner/repository"
)

type URLService interface {
	CreateShortURL(ctx context.Context, userID int64, original, short string) (*models.URL, error)
	GetByShort(ctx context.Context, short string) (*models.URL, error)
	GetByUserID(ctx context.Context, userID int64) ([]*models.URL, error)
}

type urlService struct {
	urlRepo repository.URLRepository
}

// CreateShortURL implements URLService.
func (s *urlService) CreateShortURL(ctx context.Context, userID int64, original string, short string) (*models.URL, error) {
	existing, _ := s.urlRepo.FindByShort(ctx, short)
	if existing != nil {
		return nil, errors.New("url already exists")
	}
	url := &models.URL{
		UserID:   userID,
		Original: original,
		Short:    short,
	}
	if err := s.urlRepo.Create(ctx, url); err != nil {
		return nil, err
	}
	return url, nil
}

// GetByShort implements URLService.
func (s *urlService) GetByShort(ctx context.Context, short string) (*models.URL, error) {
	return s.urlRepo.FindByShort(ctx, short)
}

// GetByUserID implements URLService.
func (s *urlService) GetByUserID(ctx context.Context, userID int64) ([]*models.URL, error) {
	return s.urlRepo.FindByUser(ctx, userID)
}

func NewURLService(urlRepo repository.URLRepository) URLService {
	return &urlService{urlRepo: urlRepo}
}
