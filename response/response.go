package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Status{
		Code:    200,
		Message: "成功",
		Data:    data,
	})
}

func Fail(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Status{
		Code:    400,
		Message: message,
		Data:    nil,
	})
}
