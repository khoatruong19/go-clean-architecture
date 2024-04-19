package todo

import (
	"github.com/gin-gonic/gin"
)

func TodoRoutes(rg *gin.RouterGroup, handler TodoHandler) {
	todos := rg.Group("/todos")

	todos.POST("/", handler.CreateTodo)
}
