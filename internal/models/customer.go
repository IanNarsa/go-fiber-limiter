package models

type Customer struct {
	CustomerId   int    `db:"customer_id"`
	CustomerName string `db:"customer_name"`
	ContactInfo  string `db:"contact_info"`
}
