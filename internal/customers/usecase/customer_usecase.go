package usecase

import (
	"context"
	"romeo-lima/internal/customers/repository"
	"romeo-lima/internal/models"
)

type CustomerUsecase interface {
	GetCustomerByID(ctx context.Context, id int) (*models.Customer, error)
	CreateCustomer(ctx context.Context, customer *models.Customer) error
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerUsecase(cr repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{cr}
}

func (cu *customerUsecase) GetCustomerByID(ctx context.Context, id int) (*models.Customer, error) {
	return cu.customerRepo.GetByID(ctx, id)
}

func (cu *customerUsecase) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	return cu.customerRepo.Store(ctx, customer)
}
