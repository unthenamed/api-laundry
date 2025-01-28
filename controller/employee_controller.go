package controller

import (
	"api-laundry/model"
	"api-laundry/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	RouterGroup *gin.RouterGroup
	Service     service.EmployeeService
}

func (c *EmployeeController) Route() {
	c.RouterGroup.POST("/employees", c.InsertEmployee)
	c.RouterGroup.GET("/employees/:id", c.GetEmployeeById)
	c.RouterGroup.GET("/employees", c.GetAllEmployees)
	c.RouterGroup.PUT("/employees/:id", c.UpdateEmployeeById)
	c.RouterGroup.DELETE("/employees/:id", c.DeleteEmployeeById)
}

func (p *EmployeeController) InsertEmployee(c *gin.Context) {
	var employee model.Employees
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	rEmployee, err := p.Service.InsertEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, rEmployee)
}

func (p *EmployeeController) GetEmployeeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.Service.DeleteEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})

}

func (p *EmployeeController) GetAllEmployees(c *gin.Context) {
	employee, err := p.Service.GetAllEmployee()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (p *EmployeeController) UpdateEmployeeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var employee model.Employees
	err = c.ShouldBind(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	employee, err = p.Service.UpdateEmployeeById(id, employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, employee)

}

func (p *EmployeeController) DeleteEmployeeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.Service.DeleteEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})

}

func ObjEmployeeController(rg *gin.RouterGroup, service service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		RouterGroup: rg,
		Service:     service,
	}
}
