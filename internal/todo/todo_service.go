package todo

import (
	"context"
	"time"

	db "github.com/khoatruong19/go-clean-architecture/db/sqlc"
)

type TodoService struct {
	repo    Repository
	timeout time.Duration
}

type NewTodoServiceParams struct {
	Repo Repository
}

func NewTodoService(params NewTodoServiceParams) Service {
	return &TodoService{
		repo:    params.Repo,
		timeout: time.Duration(2) * time.Second,
	}
}

func (s *TodoService) CreateTodo(c context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	todo, err := s.repo.CreateTodo(ctx, db.CreateTodoParams{
		UserID: req.UserID,
		Title:  req.Title,
	})
	if err != nil {
		return nil, err
	}

	res := &CreateTodoResponse{
		ID:    todo.ID,
		Title: todo.Title,
	}

	return res, nil
}
