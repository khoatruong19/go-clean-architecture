package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khoatruong19/go-clean-architecture/config"
	"github.com/khoatruong19/go-clean-architecture/initializers"
	"github.com/khoatruong19/go-clean-architecture/internal/todo"
)

var router = gin.Default()

func main() {
	cfg := config.MustLoadConfig()
	store := initializers.ConnectToDB(cfg)

	todoModule := todo.NewTodoModule(todo.NewTodoModuleParams{
		DB: store.Queries,
	})

	v1 := router.Group("/v1")
	todo.TodoRoutes(v1, *todoModule.Handler)

	router.Run(":8080")
}
