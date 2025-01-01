package repo

import (
	"api-laundry/model"
	"database/sql"
)

type EmployeeRepo interface {
	InsertEmployee(model.Employees) (model.Employees, error)
	GetEmployeeById(int) (model.Employees, error)
	GetAllEmployee() ([]model.Employees, error)
	UpdateEmployeeById(int, model.Employees) (model.Employees, error)
	DeleteEmployeeById(int) error
}

type employeeRepo struct {
	db *sql.DB
}

func (e *employeeRepo) InsertEmployee(mEmployee model.Employees) (model.Employees, error) {
	tx, err := e.db.Begin()
	if err != nil {
		return model.Employees{}, err
	}

	statement := `
		INSERT INTO employee
		(name, phone, address)
		VALUES($1, $2, $3) RETURNING id;`

	err = tx.QueryRow(
		statement,
		mEmployee.Name,
		mEmployee.PhoneNumber,
		mEmployee.Address,
	).Scan(&mEmployee.Id)

	if err != nil {
		tx.Rollback()
		return model.Employees{}, err
	}

	return mEmployee, tx.Commit()
}

func (e *employeeRepo) GetEmployeeById(id int) (model.Employees, error) {
	var mEmployee model.Employees

	statement := `SELECT id, name, phone, address FROM employee WHERE id = $1;`
	err := e.db.QueryRow(statement, id).Scan(
		&mEmployee.Id,
		&mEmployee.Name,
		&mEmployee.PhoneNumber,
		&mEmployee.Address,
	)

	if err != nil {
		return model.Employees{}, err
	}

	return mEmployee, nil
}

func (e *employeeRepo) GetAllEmployee() ([]model.Employees, error) {
	var employees []model.Employees

	statement := `SELECT id, name, phone, address FROM employee;`
	rows, err := e.db.Query(statement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee model.Employees
		err = rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.PhoneNumber,
			&employee.Address,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (e *employeeRepo) UpdateEmployeeById(id int, mEmployee model.Employees) (model.Employees, error) {
	tx, err := e.db.Begin()
	if err != nil {
		return model.Employees{}, err
	}

	statement := `UPDATE employee SET name=$1, phone=$2, address=$3 WHERE id=$4 RETURNING id;`
	err = tx.QueryRow(statement, mEmployee.Name, mEmployee.PhoneNumber, mEmployee.Address, id).Scan(&mEmployee.Id)
	if err != nil {
		tx.Rollback()
		return model.Employees{}, err
	}

	return mEmployee, tx.Commit()
}

func (e *employeeRepo) DeleteEmployeeById(id int) error {
	tx, err := e.db.Begin()
	if err != nil {
		return err
	}

	statement := `DELETE FROM employee WHERE id=$1;`
	_, err = tx.Exec(statement, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func ObjEmployeeRepo(db *sql.DB) EmployeeRepo {
	return &employeeRepo{db}
}
