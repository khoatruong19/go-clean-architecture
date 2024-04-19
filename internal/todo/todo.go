package todo

import (
	"context"

	db "github.com/khoatruong19/go-clean-architecture/db/sqlc"
)

type Repository interface {
	CreateTodo(ctx context.Context, arg db.CreateTodoParams) (db.Todo, error)
}

type Service interface {
	CreateTodo(c context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error)
}

type CreateTodoRequest struct {
	UserID int64  `json:"userId"`
	Title  string `json:"title"`
}

type CreateTodoResponse struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
