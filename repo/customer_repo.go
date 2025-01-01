package repo

import (
	"api-laundry/model"
	"database/sql"
)

type CustomerRepo interface {
	InsertCustomer(model.Customers) (model.Customers, error)
	GetCustomerById(int) (model.Customers, error)
	GetAllCustomer() ([]model.Customers, error)
	UpdateCustomerById(int, model.Customers) (model.Customers, error)
	DeleteCustomerById(int) error
}

type Customer struct {
	db *sql.DB
}

func (c *Customer) InsertCustomer(mCustomer model.Customers) (model.Customers, error) {
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

func (c *Customer) GetCustomerById(id int) (model.Customers, error) {
	var mCustomer model.Customers

	statement := `SELECT id, name, phone, address FROM customers WHERE id = $1;`
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

func (c *Customer) GetAllCustomer() ([]model.Customers, error) {
	var mCustomer []model.Customers

	statement := `SELECT id, name, phone, address FROM customers;`
	rows, err := c.db.Query(statement)
	if err != nil {
		return []model.Customers{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer model.Customers
		err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.PhoneNumber,
			&customer.Address,
		)

		if err != nil {
			return []model.Customers{}, err
		}

		mCustomer = append(mCustomer, customer)
	}

	return mCustomer, nil
}

func (c *Customer) UpdateCustomerById(id int, mCustomer model.Customers) (model.Customers, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return model.Customers{}, err
	}

	statement := `
		UPDATE customers
		SET name = $1, phone = $2, address = $3
		WHERE id = $4;`

	_, err = tx.Exec(
		statement,
		mCustomer.Name,
		mCustomer.PhoneNumber,
		mCustomer.Address,
		id,
	)

	if err != nil {
		tx.Rollback()
		return model.Customers{}, err
	}

	return mCustomer, tx.Commit()
}

func (c *Customer) DeleteCustomerById(id int) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	statement := `DELETE FROM customers WHERE id = $1;`
	_, err = tx.Exec(statement, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func ObjCustomerRepo(db *sql.DB) CustomerRepo {
	return &Customer{
		db: db,
	}
}
