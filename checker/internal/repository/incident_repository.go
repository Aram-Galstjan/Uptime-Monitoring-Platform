package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"checker/internal/models"
)

type IncidentRepository struct {
	db *pgxpool.Pool
}

// NewIncidentRepository создает новый репозиторий.
func NewIncidentRepository(db *pgxpool.Pool) *IncidentRepository {
	return &IncidentRepository{
		db: db,
	}
}

// Create создает новый инцидент.
func (r *IncidentRepository) Create(ctx context.Context, incident models.Incident) error {
	// Сохранение нового инцидента в базу данных.
	return nil
}

// Close закрывает активный инцидент.
func (r *IncidentRepository) Close(ctx context.Context, siteID int) error {
	// Обновление инцидента: установка времени окончания и отметка, что он завершен.
	return nil
}

// GetActive возвращает активный инцидент для указанного сайта.
func (r *IncidentRepository) GetActive(ctx context.Context, siteID int) (*models.Incident, error) {
	// Получение активного инцидента из базы данных.
	return nil, nil
}
