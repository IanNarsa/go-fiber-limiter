package repository

import (
	"context"
	"romeo-lima/internal/models"
)

type CustomerRepository interface {
	GetByID(ctx context.Context, id int) (*models.Customer, error)
	Store(ctx context.Context, customer *models.Customer) error
}
