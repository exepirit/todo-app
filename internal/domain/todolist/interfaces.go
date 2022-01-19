package todolist

import (
	"context"
	"github.com/google/uuid"
)

// IRepository provides TODO lists storage.
type IRepository interface {
	GetByUserID(ctx context.Context, userId uuid.UUID) ([]*TodoList, error)
	GetByID(ctx context.Context, id uuid.UUID) (*TodoList, error)
	Put(ctx context.Context, todolist *TodoList) error
	Update(ctx context.Context, todolist *TodoList) error
}
