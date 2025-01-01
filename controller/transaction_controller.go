package controller

import (
	"api-laundry/model"
	"api-laundry/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	RouterGroup *gin.RouterGroup
	Service     service.TransactionService
}

func (controller *TransactionController) Route() {
	controller.RouterGroup.GET("/transactions", controller.GetAllTransactions)
	controller.RouterGroup.GET("/transactions/:id", controller.GetTransactionById)
	controller.RouterGroup.POST("/transactions", controller.InsertTransaction)
}

func (controller *TransactionController) GetAllTransactions(c *gin.Context) {
	model := model.Transaction{}
	model.Query.StartDate = c.Query("startDate")
	model.Query.EndDate = c.Query("endDate")
	model.Query.ProductName = c.Query("productName")

	transaction, err := controller.Service.GetAllTransaction(model)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, transaction)
}

func (controller *TransactionController) GetTransactionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	transaction, err := controller.Service.GetTransactionById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, transaction.Response)
}

func (controller *TransactionController) InsertTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	transaction, err := controller.Service.InsertTransaction(transaction)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, transaction.Response)
}

func ObjTransactionController(router *gin.RouterGroup, service service.TransactionService) *TransactionController {
	return &TransactionController{
		RouterGroup: router,
		Service:     service,
	}
}
