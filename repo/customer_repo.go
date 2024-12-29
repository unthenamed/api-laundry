package repo

import (
	"api-laundry/model"
	"database/sql"
)

type CustomerRepo interface {
	insertCustomer(model.Customers) (model.Customers, error)
	getCustomerById(int) (model.Customers, error)
	getAllCustomer() ([]model.Customers, error)
}

type Customer struct {
	db *sql.DB
}

func (c *Customer) insertCustomer(mCustomer model.Customers) (model.Customers, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return model.Customers{}, err
	}

	statement := `
		INSERT INTO customers
		(name, phone, address)
		VALUES($1, $2, $3) RETURNING id;`

	err = tx.QueryRow(
		statement,
		mCustomer.Name,
		mCustomer.PhoneNumber,
		mCustomer.Address,
	).Scan(&mCustomer.Id)

	if err != nil {
		tx.Rollback()
		return model.Customers{}, err
	}

	return mCustomer, tx.Commit()
}

func (c *Customer) getCustomerById(id int) (model.Customers, error) {
	var mCustomer model.Customers

	statement := `SELECT * FROM customers WHERE id = $1;`
	err := c.db.QueryRow(statement, id).Scan(
		&mCustomer.Id,
		&mCustomer.Name,
		&mCustomer.PhoneNumber,
		&mCustomer.Address,
	)

	if err != nil {
		return model.Customers{}, err
	}

	return mCustomer, nil
}

func (c *Customer) getAllCustomer() ([]model.Customers, error) {
	var nCustomers []model.Customers

	return nCustomers, nil
}

func ObjCustomerRepo(db *sql.DB) CustomerRepo {
	return &Customer{
		db: db,
	}
}
