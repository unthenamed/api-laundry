package controller

import (
	"api-laundry/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	RouterGroup *gin.RouterGroup
	Service     service.ProductService
}

func (p *ProductController) Route() {

}

func ObjProductController(rg *gin.RouterGroup, service service.ProductService) *ProductController {
	return &ProductController{
		RouterGroup: rg,
		Service:     service,
	}
}
