package controller

import (
	"api-laundry/model"
	"api-laundry/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	RouterGroup *gin.RouterGroup
	Service     service.ProductService
}

func (c *ProductController) Route() {
	c.RouterGroup.POST("/products", c.InsertProduct)
	c.RouterGroup.GET("/products/:id", c.GetProductById)
	c.RouterGroup.GET("/products", c.GetAllProducts)
	c.RouterGroup.PUT("/products/:id", c.UpdateProductById)
	c.RouterGroup.DELETE("/products/:id", c.DeleteProductById)
}

func (p *ProductController) InsertProduct(c *gin.Context) {
	var product model.Products
	err := c.ShouldBind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	product, err = p.Service.InsertProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, product)

}

func (p *ProductController) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	product, err := p.Service.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *ProductController) GetAllProducts(c *gin.Context) {
	productName := c.Query("productName")

	product, err := p.Service.GetAllProduct(productName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *ProductController) UpdateProductById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var product model.Products
	err = c.ShouldBind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	product, err = p.Service.UpdateProductById(id, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *ProductController) DeleteProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.Service.DeleteProductById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})

}
func ObjProductController(rg *gin.RouterGroup, service service.ProductService) *ProductController {
	return &ProductController{
		RouterGroup: rg,
		Service:     service,
	}
}
