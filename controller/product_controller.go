package controller

import (
	"api-laundry/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) InsertProduct(c *gin.Context) {
	var product model.Products
	err := c.ShouldBind(&product)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	product, err = h.Service.InsertProduct(product)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "created", product)

}

func (h *Handlers) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	product, err := h.Service.GetProductById(id)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "OK", product)
}

func (h *Handlers) GetAllProducts(c *gin.Context) {
	productName := c.Query("productName")

	product, err := h.Service.GetAllProduct(productName)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "OK", product)
}

func (h *Handlers) UpdateProductById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	var product model.Products
	err = c.ShouldBind(&product)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	product, err = h.Service.UpdateProductById(id, product)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "OK", product)
}

func (h *Handlers) DeleteProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	err = h.Service.DeleteProductById(id)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})

}
