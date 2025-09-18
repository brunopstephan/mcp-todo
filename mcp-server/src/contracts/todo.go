package contracts

import "mcp-server/src/schemas"

type TodoContract interface {
	GetTodos() ([]schemas.Todo, error)
	CreateTodo(todo schemas.Todo) error
	UpdateTodo(id string, todo schemas.Todo) error
	DeleteTodo(id string) error
	MultiCreateTodo(todos []schemas.Todo) error
}
