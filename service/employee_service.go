package service

import (
	"api-laundry/model"
	"fmt"
)

func (s *laundryService) InsertEmployee(mEmployee model.Employees) (model.Employees, error) {
	return s.employees.InsertEmployee(mEmployee)
}

func (s *laundryService) GetEmployeeById(id int) (model.Employees, error) {
	return s.employees.GetEmployeeById(id)
}

func (s *laundryService) GetAllEmployee() ([]model.Employees, error) {
	return s.employees.GetAllEmployee()

}

func (s *laundryService) UpdateEmployeeById(id int, mEmployee model.Employees) (model.Employees, error) {
	oldEmployee, err := s.employees.GetEmployeeById(id)
	if err != nil {
		return model.Employees{}, err
	}

	if mEmployee.Name == "" {
		mEmployee.Name = oldEmployee.Name
	}
	if mEmployee.Address == "" {
		mEmployee.Address = oldEmployee.Address
	}
	if mEmployee.PhoneNumber == "" {
		mEmployee.PhoneNumber = oldEmployee.PhoneNumber
	}

	return s.employees.UpdateEmployeeById(id, mEmployee)
}

func (s *laundryService) DeleteEmployeeById(id int) error {
	_, err := s.employees.GetEmployeeById(id)
	if err != nil {
		return err
	}

	trx, err := s.transactions.GetAllTransaction(model.Transaction{})
	for _, t := range trx {
		if t.Employee.Id == id {
			return fmt.Errorf("Id Already in transaction, pleas delete transaction first!")
		}
	}

	return s.employees.DeleteEmployeeById(id)
}
