package controller

import (
	"api-laundry/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	RouterGroup *gin.RouterGroup
	Service     service.EmployeeService
}

func (controller *EmployeeController) Route() {

}

func ObjEmployeeController(rg *gin.RouterGroup, service service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		RouterGroup: rg,
		Service:     service,
	}
}
