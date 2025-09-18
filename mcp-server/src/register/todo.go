package register

import (
	"mcp-server/src/repositories"
	"mcp-server/src/tools"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"gorm.io/gorm"
)

func RegisterTodoTools(server *mcp.Server, db *gorm.DB) {

	todosRepo := repositories.NewTodoRepository(db)

	todosTools := tools.NewTodoTools(todosRepo)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get-todos",
		Description: "Search for all to-dos",
	}, todosTools.GetTodos)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create-todo",
		Description: "Create a todo with a title and description",
	}, todosTools.CreateTodo)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "multi-create-todo",
		Description: "Create multiple todos with a title and description",
	}, todosTools.MultiCreateTodo)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update-todo",
		Description: "Update an existing todo",
	}, todosTools.UpdateTodo)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete-todos",
		Description: "Delete an existing todo",
	}, todosTools.DeleteTodo)

}
