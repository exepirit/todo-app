package repository

import (
	"context"

	domain "github.com/exepirit/todo-app/internal/domain/todolist"

	"github.com/google/uuid"
)

// ITodoListRepo provides TODO lists storage.
type ITodoListRepo interface {
	GetByUserID(ctx context.Context, userId uuid.UUID) (*domain.TodoList, error)
	Put(ctx context.Context, todolist *domain.TodoList) error
	Update(ctx context.Context, todolist *domain.TodoList) error
}
