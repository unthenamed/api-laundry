package service

import (
	"api-laundry/model"
	"api-laundry/repo"
)

type CustomerService interface {
	InsertCustomer(model.Customers) (model.Customers, error)
	GetAllCustomer() ([]model.Customers, error)
	GetCustomerById(int) (model.Customers, error)
	UpdateCustomerById(int, model.Customers) (model.Customers, error)
	DeleteCustomerById(int) error
}

type customerService struct {
	repo repo.CustomerRepo
}

func (c *customerService) InsertCustomer(mCustomer model.Customers) (model.Customers, error) {
	return c.repo.InsertCustomer(mCustomer)
}

func (c *customerService) GetAllCustomer() ([]model.Customers, error) {
	return c.repo.GetAllCustomer()
}

func (c *customerService) GetCustomerById(id int) (model.Customers, error) {
	return c.repo.GetCustomerById(id)
}

func (c *customerService) UpdateCustomerById(id int, mCustomer model.Customers) (model.Customers, error) {

	oldCustomer, err := c.repo.GetCustomerById(id)
	if err != nil {
		return model.Customers{}, err
	}

	if mCustomer.Name == "" {
		mCustomer.Name = oldCustomer.Name
	}

	if mCustomer.PhoneNumber == "" {
		mCustomer.PhoneNumber = oldCustomer.PhoneNumber
	}

	if mCustomer.Address == "" {
		mCustomer.Address = oldCustomer.Address
	}

	return c.repo.UpdateCustomerById(id, mCustomer)
}

func (c *customerService) DeleteCustomerById(id int) error {
	return c.repo.DeleteCustomerById(id)
}

func ObjCustomerService(repo repo.CustomerRepo) CustomerService {
	return &customerService{repo}
}
