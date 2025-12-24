package repository

import (
	"github.com/Rokli/URL-Shortener/internal/domain"
)

type UrlRepository interface {
	Create(url *domain.Url)
	FindByShortCode(code string) (*domain.Url, error)
	IncrementVisits(code string) error
	Delete(code string) error
}
