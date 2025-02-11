package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type message struct {
	Message string
	Data    any
}

type httpRespon struct {
	Code int
	message
}

var notFoundError string = "no rows"
var badRequestError string = "parsing"
var notModifedError string = "constrain"

func strValidation(s, ss string) bool {
	return strings.Contains(s, strings.TrimSpace(ss))
}

func ErrorRespon(c *gin.Context, err error) {
	respon := httpRespon{
		Code:    http.StatusInternalServerError,
		message: message{Message: err.Error()},
	}

	if strValidation(err.Error(), notFoundError) {
		respon.message.Message = "Not Found!"
		respon.Code = http.StatusNotFound
	}

	if strValidation(err.Error(), badRequestError) {
		respon.message.Message = "Bad Request!"
		respon.Code = http.StatusBadRequest
	}

	if strValidation(err.Error(), notModifedError) {
		respon.message.Message = "Already in transaction!"
		respon.Code = http.StatusBadRequest
	}

	c.JSON(respon.Code, gin.H{"message": respon.message.Message})

}

func SuccesRespon(c *gin.Context, msg string, data any) {
	respon := httpRespon{
		Code: http.StatusOK,
		message: message{
			Message: msg,
			Data:    data,
		},
	}

	if msg == "created" {
		respon.Code = http.StatusCreated
	}

	c.JSON(respon.Code, respon.message)
}
