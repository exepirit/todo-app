package repository_test

import (
	"context"
	"testing"

	domain "github.com/exepirit/todo-app/internal/domain/todolist"
	"github.com/exepirit/todo-app/internal/repository"
)

func TestInMemoryTodoListRepository(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewMemoryTodoList()
	todolist := domain.TodoList{}

	if err := repo.Put(ctx, &todolist); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	_, err := repo.GetByID(ctx, todolist.ID())
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
