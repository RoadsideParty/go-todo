package routers

import (
	"github.com/gin-gonic/gin"
	"todo/controller"
)

func SetupRouter(r *gin.Engine) {
	todoGroup := r.Group("v1")
	todoGroup.POST("/todo", controller.CreateTodo)
}
