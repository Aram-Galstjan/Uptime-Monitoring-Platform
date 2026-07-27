package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"checker/internal/models"
)

type SiteRepository struct {
	db *pgxpool.Pool
}

// NewSiteRepository создает новый репозиторий.
func NewSiteRepository(db *pgxpool.Pool) *SiteRepository {
	return &SiteRepository{
		db: db,
	}
}

// GetAll возвращает список всех сайтов.
func (r *SiteRepository) GetAll(ctx context.Context) ([]models.Site, error) {
	// Получение списка сайтов из базы данных.
	return nil, nil
}

// Create добавляет новый сайт.
func (r *SiteRepository) Create(ctx context.Context, site models.Site) error {
	// Добавление сайта в базу данных.
	return nil
}

// Update изменяет информацию о сайте.
func (r *SiteRepository) Update(ctx context.Context, site models.Site) error {
	// Обновление данных сайта в базе данных.
	return nil
}

// Delete удаляет сайт по ID.
func (r *SiteRepository) Delete(ctx context.Context, id int) error {
	// Удаление сайта из базы данных.
	return nil
}
