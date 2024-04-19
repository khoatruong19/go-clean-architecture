package todo

import db "github.com/khoatruong19/go-clean-architecture/db/sqlc"

type TodoServices struct {
	TodoService Service
}

type TodoModule struct {
	Handler  *TodoHandler
	Services TodoServices
}

type NewTodoModuleParams struct {
	DB *db.Queries
}

func NewTodoModule(params NewTodoModuleParams) *TodoModule {
	todoSvc := NewTodoService(NewTodoServiceParams{
		Repo: params.DB,
	})

	return &TodoModule{
		Handler: NewTodoHandler(NewTodoHandlerParams{
			Svc: todoSvc,
		}),
		Services: TodoServices{
			TodoService: todoSvc,
		},
	}
}
