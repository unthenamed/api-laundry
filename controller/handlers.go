package controller

import (
	"api-laundry/service"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	RouterGroup *gin.RouterGroup
	Service     service.LaundryService
}

func (h *Handlers) Route() {

	// Transaction Endpoint
	h.RouterGroup.GET("/transactions", h.GetAllTransactions)
	h.RouterGroup.GET("/transactions/:id", h.GetTransactionById)
	h.RouterGroup.POST("/transactions", h.InsertTransaction)

	// Customer Endpoint
	h.RouterGroup.GET("/customers", h.GetAllCustomers)
	h.RouterGroup.GET("/customers/:id", h.GetCustomerById)
	h.RouterGroup.POST("/customers", h.InsertCustomer)
	h.RouterGroup.PUT("/customers/:id", h.UpdateCustomerById)
	h.RouterGroup.DELETE("/customers/:id", h.DeleteCustomerById)

	// Employee Endpoint
	h.RouterGroup.POST("/employees", h.InsertEmployee)
	h.RouterGroup.GET("/employees/:id", h.GetEmployeeById)
	h.RouterGroup.GET("/employees", h.GetAllEmployees)
	h.RouterGroup.PUT("/employees/:id", h.UpdateEmployeeById)
	h.RouterGroup.DELETE("/employees/:id", h.DeleteEmployeeById)

	// Product Endpoint
	h.RouterGroup.POST("/products", h.InsertProduct)
	h.RouterGroup.GET("/products/:id", h.GetProductById)
	h.RouterGroup.GET("/products", h.GetAllProducts)
	h.RouterGroup.PUT("/products/:id", h.UpdateProductById)
	h.RouterGroup.DELETE("/products/:id", h.DeleteProductById)

}

func NewHandlersController(routerGroup *gin.RouterGroup, service service.LaundryService) *Handlers {
	return &Handlers{
		RouterGroup: routerGroup,
		Service:     service,
	}
}
