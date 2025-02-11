package controller

import (
	"api-laundry/model"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetAllTransactions(c *gin.Context) {
	model := model.Transaction{}
	model.Query.StartDate = c.Query("startDate")
	model.Query.EndDate = c.Query("endDate")
	model.Query.ProductName = c.Query("productName")

	transaction, err := h.Service.GetAllTransaction(model)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	if len(transaction) == 0 {
		ErrorRespon(c, errors.New("no rows"))
		return
	}

	SuccesRespon(c, "List all transaction ", transaction)
}

func (h *Handlers) GetTransactionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	transaction, err := h.Service.GetTransactionById(id)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, fmt.Sprintf("Transaction by id %d", id), transaction.Response)
}

func (h *Handlers) InsertTransaction(c *gin.Context) {
	var transaction model.Transaction
	err := c.ShouldBind(&transaction.Bills)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	rTransaction, err := h.Service.InsertTransaction(transaction)
	if err != nil {
		ErrorRespon(c, err)
		return
	}

	SuccesRespon(c, "New transaction created!", rTransaction)

}
