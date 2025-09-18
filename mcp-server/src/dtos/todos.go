package dtos

import "mcp-server/src/schemas"

type Empty struct{}

type Default struct {
	Message string `json:"message"`
}

type GetTodosOutput struct {
	Todos []schemas.Todo `json:"todos"`
}

type DeleteTodoInput struct {
	ID string `json:"id"`
}

type MultiCreateTodoInput struct {
	Todos []schemas.Todo `json:"todos"`
}

type UpdateTodoInput struct {
	ID   string       `json:"id"`
	Todo schemas.Todo `json:"todo"`
}
