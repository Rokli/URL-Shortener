package repository

import (
	"database/sql"

	"github.com/Rokli/URL-Shortener/internal/domain"
)

type UrlRepository interface {
	Create(url *domain.Url) error
	Get(id int) *domain.Url
	Update(id int) *domain.Url
	FindByShortCode(code string) (*domain.Url, error)
	IncrementVisits(code string) error
	Delete(code string) error
}

type postgresUrlRepository struct {
	db *sql.DB
}

func NewPostgresUrlRepository(db *sql.DB) UrlRepository {
	return &postgresUrlRepository{db: db}
}

func (r *postgresUrlRepository) Get(id int) *domain.Url {
	return domain.CreateNewUrl()
}

func (r *postgresUrlRepository) Update(id int) *domain.Url {
	return domain.CreateNewUrl()
}

func (r *postgresUrlRepository) Create(url *domain.Url) error {
	return nil
}

func (r *postgresUrlRepository) FindByShortCode(code string) (*domain.Url, error) {
	return domain.CreateNewUrl(), nil
}

func (r *postgresUrlRepository) IncrementVisits(code string) error {
	return nil
}

func (r *postgresUrlRepository) Delete(code string) error {
	return nil
}
