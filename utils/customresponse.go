package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(c *gin.Context, data interface{}) {
	response := CustomResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c *gin.Context, statusCode int) {
	var message string

	if statusCode == 404 {
		message = "Data Not Found."
	}

	if statusCode == 400 {
		message = "Operation Failed."
	}

	if statusCode == 500 {
		message = "Operation Failed."
	}

	response := CustomResponse{
		Status:  statusCode,
		Message: message,
	}

	c.JSON(statusCode, response)
}
