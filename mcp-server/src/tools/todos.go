package tools

import (
	"context"
	"mcp-server/src/contracts"
	"mcp-server/src/dtos"
	"mcp-server/src/schemas"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type TodoTools struct {
	toolsRepository contracts.TodoContract
}

func NewTodoTools(db contracts.TodoContract) *TodoTools {
	return &TodoTools{toolsRepository: db}
}

func (t *TodoTools) GetTodos(ctx context.Context, req *mcp.CallToolRequest, input dtos.Empty) (*mcp.CallToolResult, dtos.GetTodosOutput, error) {
	todos, err := t.toolsRepository.GetTodos()
	if err != nil {
		return nil, dtos.GetTodosOutput{}, err
	}

	return nil, dtos.GetTodosOutput{Todos: todos}, nil
}

func (t *TodoTools) CreateTodo(ctx context.Context, req *mcp.CallToolRequest, input schemas.Todo) (*mcp.CallToolResult, dtos.Default, error) {
	todo := schemas.Todo{
		Name:        input.Name,
		Description: input.Description,
		Done:        input.Done,
		Date:        input.Date,
	}

	err := t.toolsRepository.CreateTodo(todo)
	if err != nil {
		return nil, dtos.Default{}, err
	}

	return nil, dtos.Default{Message: "Todo created successfully"}, nil
}

func (t *TodoTools) MultiCreateTodo(ctx context.Context, req *mcp.CallToolRequest, input dtos.MultiCreateTodoInput) (*mcp.CallToolResult, dtos.Default, error) {

	err := t.toolsRepository.MultiCreateTodo(input.Todos)
	if err != nil {
		return nil, dtos.Default{}, err
	}

	return nil, dtos.Default{
		Message: "Todos created successfully",
	}, nil
}

func (t *TodoTools) UpdateTodo(ctx context.Context, req *mcp.CallToolRequest, input dtos.UpdateTodoInput) (*mcp.CallToolResult, dtos.Default, error) {

	err := t.toolsRepository.UpdateTodo(input.ID, input.Todo)
	if err != nil {
		return nil, dtos.Default{}, err
	}

	return nil, dtos.Default{
		Message: "Todo updated successfully",
	}, nil
}

func (t *TodoTools) DeleteTodo(ctx context.Context, req *mcp.CallToolRequest, input dtos.DeleteTodoInput) (*mcp.CallToolResult, dtos.Default, error) {

	err := t.toolsRepository.DeleteTodo(input.ID)
	if err != nil {
		return nil, dtos.Default{}, err
	}

	return nil, dtos.Default{
		Message: "Todo deleted successfully",
	}, nil
}
