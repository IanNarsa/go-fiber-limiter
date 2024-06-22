package repository

import (
	"context"
	"database/sql"
	"romeo-lima/internal/models"

	"github.com/jmoiron/sqlx"
)

type postgresCustomerRepository struct {
	Conn *sqlx.DB
}

func NewPostgresCustomerRepository(conn *sqlx.DB) CustomerRepository {
	return &postgresCustomerRepository{conn}
}

func (p *postgresCustomerRepository) GetAllData(ctx context.Context) (*[]models.Customer, error) {
	query := "SELECT customer_id, customer_name, contact_info FROM customers"
	var customer []models.Customer
	err := p.Conn.SelectContext(ctx, &customer, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (p *postgresCustomerRepository) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	query := "SELECT customer_id, customer_name, contact_info FROM customers WHERE customer_id = $1"
	var customer models.Customer
	err := p.Conn.GetContext(ctx, &customer, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (p *postgresCustomerRepository) Store(ctx context.Context, customer *models.Customer) error {
	query := "INSERT INTO customers (name, email) VALUES ($1, $2)"
	_, err := p.Conn.ExecContext(ctx, query, customer.CustomerName, customer.ContactInfo)
	return err
}
