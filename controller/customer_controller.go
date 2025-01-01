package controller

import (
	"api-laundry/model"
	"api-laundry/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	RouterGroup *gin.RouterGroup
	Service     service.CustomerService
}

func (controller *CustomerController) Route() {
	controller.RouterGroup.GET("/customers", controller.GetAllCustomers)
	controller.RouterGroup.GET("/customers/:id", controller.GetCustomerById)
	controller.RouterGroup.POST("/customers", controller.InsertCustomer)
	controller.RouterGroup.PUT("/customers/:id", controller.UpdateCustomerById)
	controller.RouterGroup.DELETE("/customers/:id", controller.DeleteCustomerById)
}

func (controller *CustomerController) GetAllCustomers(c *gin.Context) {
	customers, err := controller.Service.GetAllCustomer()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, customers)
}

func (controller *CustomerController) GetCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	customer, err := controller.Service.GetCustomerById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, customer)
}

func (controller *CustomerController) InsertCustomer(c *gin.Context) {
	var customer model.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	customer, err := controller.Service.InsertCustomer(customer)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, customer)
}

func (controller *CustomerController) UpdateCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	var customer model.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	customer, err = controller.Service.UpdateCustomerById(id, customer)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, customer)
}

func (controller *CustomerController) DeleteCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = controller.Service.DeleteCustomerById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Customer deleted",
	})
}

func ObjCustomerController(router *gin.RouterGroup, service service.CustomerService) *CustomerController {
	return &CustomerController{
		RouterGroup: router,
		Service:     service,
	}
}
