package service

import (
	"api-laundry/model"
	"api-laundry/repo"
)

type EmployeeService interface {
	InsertEmployee(model.Employees) (model.Employees, error)
	GetEmployeeById(int) (model.Employees, error)
	GetAllEmployee() ([]model.Employees, error)
	UpdateEmployeeById(int, model.Employees) (model.Employees, error)
	DeleteEmployeeById(int) error
}

type employeeService struct {
	Repo repo.EmployeeRepo
}

func (e *employeeService) InsertEmployee(mEmployee model.Employees) (model.Employees, error) {
	return e.Repo.InsertEmployee(mEmployee)
}

func (e *employeeService) GetEmployeeById(id int) (model.Employees, error) {
	return e.Repo.GetEmployeeById(id)
}

func (e *employeeService) GetAllEmployee() ([]model.Employees, error) {
	return e.Repo.GetAllEmployee()

}

func (e *employeeService) UpdateEmployeeById(id int, mEmployee model.Employees) (model.Employees, error) {
	oldEmployee, err := e.Repo.GetEmployeeById(id)
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

	return e.Repo.UpdateEmployeeById(id, mEmployee)
}

func (e *employeeService) DeleteEmployeeById(id int) error {
	return e.Repo.DeleteEmployeeById(id)
}

func ObjEmployeeService(repo repo.EmployeeRepo) EmployeeService {
	return &employeeService{Repo: repo}
}
