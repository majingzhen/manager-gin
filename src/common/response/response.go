package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FailWithMessage(message string, c *gin.Context) {
	response := Response{
		Code:    1,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}

func OkWithData(data interface{}, c *gin.Context) {
	response := Response{
		Code:    0,
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

func OkWithMessage(message string, c *gin.Context) {
	response := Response{
		Code:    0,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	response := Response{
		Code:    0,
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}
