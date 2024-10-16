package utility

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

// ResponseOK is a function for send result to client
func ResponseOK(c *gin.Context, code int, obj interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: "success",
		Data:    obj,
	})
}

// ResponseError is a function for send result error to client
func ResponseError(c *gin.Context, code int, msg error, obj interface{}, errorMsg string) {
	c.JSON(code, Response{
		Code:    code,
		Message: msg.Error(),
		Data:    obj,
		Error:   errorMsg,
	})
}
