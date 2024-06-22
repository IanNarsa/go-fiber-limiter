package repository

import (
	"context"
	"romeo-lima/internal/models"
)

type CustomerRepository interface {
	GetAllData(ctx context.Context) (*[]models.Customer, error)
	GetByID(ctx context.Context, id int) (*models.Customer, error)
	Store(ctx context.Context, customer *models.Customer) error
}
