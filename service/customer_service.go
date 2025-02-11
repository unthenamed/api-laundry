package service

import (
	"api-laundry/model"
	"fmt"
)

func (s *laundryService) InsertCustomer(mCustomer model.Customers) (model.Customers, error) {
	return s.customers.InsertCustomer(mCustomer)
}

func (s *laundryService) GetAllCustomer() ([]model.Customers, error) {
	return s.customers.GetAllCustomer()
}

func (s *laundryService) GetCustomerById(id int) (model.Customers, error) {
	return s.customers.GetCustomerById(id)
}

func (s *laundryService) UpdateCustomerById(id int, mCustomer model.Customers) (model.Customers, error) {

	oldCustomer, err := s.customers.GetCustomerById(id)
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

	return s.customers.UpdateCustomerById(id, mCustomer)
}

func (s *laundryService) DeleteCustomerById(id int) error {
	_, err := s.customers.GetCustomerById(id)
	if err != nil {
		return err
	}

	trx, err := s.transactions.GetAllTransaction(model.Transaction{})
	for _, t := range trx {
		if t.Customer.Id == id {
			return fmt.Errorf("Id Already in transaction, pleas delete transaction first!")
		}
	}

	return s.customers.DeleteCustomerById(id)
}
