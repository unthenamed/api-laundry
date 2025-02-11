package controller

import (
	"api-laundry/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetAllCustomers(c *gin.Context) {
	customers, err := h.Service.GetAllCustomer()
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "OK", customers)
}

func (h *Handlers) GetCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	customer, err := h.Service.GetCustomerById(id)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "OK", customer)
}

func (h *Handlers) InsertCustomer(c *gin.Context) {
	var customer model.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		ErrorRespon(c, err)
		return
	}

	customer, err := h.Service.InsertCustomer(customer)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "created", customer)
}

func (h *Handlers) UpdateCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	var customer model.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		ErrorRespon(c, err)
		return
	}

	customer, err = h.Service.UpdateCustomerById(id, customer)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "OK", customer)
}

func (h *Handlers) DeleteCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	err = h.Service.DeleteCustomerById(id)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer deleted",
	})
}
