package repositories

import (
	"fmt"
	"mcp-server/src/contracts"
	"mcp-server/src/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) contracts.TodoContract {
	return &TodoRepository{
		db: db,
	}
}

func (t *TodoRepository) CreateTodo(todo schemas.Todo) error {
	todo.ID = uuid.New().String()
	return t.db.Create(&todo).Error
}

func (t *TodoRepository) MultiCreateTodo(todos []schemas.Todo) error {
	for i := range todos {
		todos[i].ID = uuid.New().String()
	}
	return t.db.Create(&todos).Error
}

func (t *TodoRepository) DeleteTodo(id string) error {

	result := t.db.Delete(&schemas.Todo{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("todo with id %s not found", id)
	}
	return result.Error
}

func (t *TodoRepository) GetTodos() ([]schemas.Todo, error) {
	var todos []schemas.Todo
	result := t.db.Find(&todos)
	return todos, result.Error
}

func (t *TodoRepository) UpdateTodo(id string, todo schemas.Todo) error {
	var existing schemas.Todo
	if err := t.db.First(&existing, "id = ?", id).Error; err != nil {
		return fmt.Errorf("todo with id %s not found", id)
	}

	updates := map[string]interface{}{}
	if todo.Name != "" {
		updates["name"] = todo.Name
	}
	if todo.Description != "" {
		updates["description"] = todo.Description
	}
	updates["done"] = todo.Done
	if todo.Date != "" {
		updates["date"] = todo.Date
	}

	if len(updates) == 0 {
		return nil
	}

	return t.db.Model(&existing).Updates(updates).Error
}
