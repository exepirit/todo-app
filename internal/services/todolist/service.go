package todolist

import (
	"context"

	domain "github.com/exepirit/todo-app/internal/domain/todolist"

	"github.com/google/uuid"
)

// IService provides actions related to TODO list.
type IService interface {
	// GetUserLists returns user's TODO lists slice.
	GetUserLists(ctx context.Context, userId uuid.UUID) []domain.TodoList

	// Create creates new TODO list.
	Create(ctx context.Context, userId uuid.UUID) domain.TodoList

	// PutItem puts new item into TODO list.
	PutItem(ctx context.Context, listId uuid.UUID, item domain.TodoItem)
}
