package service

import (
	"api-laundry/model"
	"api-laundry/repo"
)

type LaundryService interface {

	// Transaction
	InsertTransaction(model.Transaction) (model.Response, error)
	GetAllTransaction(model.Transaction) ([]model.Response, error)
	GetTransactionById(int) (model.Transaction, error)

	// Product
	InsertProduct(model.Products) (model.Products, error)
	GetProductById(int) (model.Products, error)
	GetAllProduct(string) ([]model.Products, error)
	UpdateProductById(int, model.Products) (model.Products, error)
	DeleteProductById(int) error

	// Customer
	InsertCustomer(model.Customers) (model.Customers, error)
	GetAllCustomer() ([]model.Customers, error)
	GetCustomerById(int) (model.Customers, error)
	UpdateCustomerById(int, model.Customers) (model.Customers, error)
	DeleteCustomerById(int) error

	// Employee
	InsertEmployee(model.Employees) (model.Employees, error)
	GetEmployeeById(int) (model.Employees, error)
	GetAllEmployee() ([]model.Employees, error)
	UpdateEmployeeById(int, model.Employees) (model.Employees, error)
	DeleteEmployeeById(int) error
}

type laundryService struct {
	transactions repo.TransactionRepo
	products     repo.ProductRepo
	customers    repo.CustomerRepo
	employees    repo.EmployeeRepo
}

/*

	fungsi yang mengimplement struct di buat terpisah
	menjadi beberapa file sesuai dengan Layanan

*/

// Constructor Service
func NewLaundryService(
	transactions repo.TransactionRepo,
	products repo.ProductRepo,
	customers repo.CustomerRepo,
	employees repo.EmployeeRepo,
) LaundryService {
	return &laundryService{
		transactions: transactions,
		products:     products,
		customers:    customers,
		employees:    employees,
	}
}
