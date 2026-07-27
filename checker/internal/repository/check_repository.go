package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"your_project/internal/models"
)

type CheckRepository struct {
	db *pgxpool.Pool
}

// NewCheckRepository создает новый репозиторий.
func NewCheckRepository(db *pgxpool.Pool) *CheckRepository {
	return &CheckRepository{
		db: db,
	}
}

// Save сохраняет результат проверки сайта.
func (r *CheckRepository) Save(ctx context.Context, result models.CheckResult) error {
	// Сохранение результата проверки в базу данных.
	return nil
}

// GetHistory возвращает историю проверок сайта.
func (r *CheckRepository) GetHistory(ctx context.Context, siteID int) ([]models.CheckResult, error) {
	// Получение истории проверок сайта из базы данных.
	return nil, nil
}
