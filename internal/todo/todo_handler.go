package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	svc Service
}

type NewTodoHandlerParams struct {
	Svc Service
}

func NewTodoHandler(params NewTodoHandlerParams) *TodoHandler {
	return &TodoHandler{
		svc: params.Svc,
	}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var t CreateTodoRequest
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.svc.CreateTodo(c.Request.Context(), &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
