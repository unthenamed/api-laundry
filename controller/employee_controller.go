package controller

import (
	"api-laundry/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) InsertEmployee(c *gin.Context) {
	var employee model.Employees
	if err := c.ShouldBindJSON(&employee); err != nil {
		ErrorRespon(c, err)
		return
	}

	rEmployee, err := h.Service.InsertEmployee(employee)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "created", rEmployee)
}

func (h *Handlers) GetEmployeeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	var nEmployee model.Employees
	nEmployee, err = h.Service.GetEmployeeById(id)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "OK", nEmployee)

}

func (h *Handlers) GetAllEmployees(c *gin.Context) {
	employee, err := h.Service.GetAllEmployee()
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "OK", employee)
}

func (h *Handlers) UpdateEmployeeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	var employee model.Employees
	err = c.ShouldBind(&employee)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	employee, err = h.Service.UpdateEmployeeById(id, employee)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "OK", employee)

}

func (h *Handlers) DeleteEmployeeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	err = h.Service.DeleteEmployeeById(id)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})

}
