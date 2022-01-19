package todolist

import (
	"context"
	"fmt"

	domain "github.com/exepirit/todo-app/internal/domain/todolist"
	"github.com/google/uuid"
)

// IService provides actions related to TODO list.
type IService interface {
	// GetUserLists returns user's TODO lists slice.
	GetUserLists(ctx context.Context, userId uuid.UUID) ([]*domain.TodoList, error)

	// Create creates new TODO list.
	Create(ctx context.Context, user domain.User) (*domain.TodoList, error)

	// PutItem puts new item into TODO list.
	PutItem(ctx context.Context, listId uuid.UUID, item *domain.TodoItem) error
}

// NewService creates new lists service.
func NewService(todoListRepo domain.IRepository) IService {
	return &service{
		todoLists: todoListRepo,
		factory:   domain.NewTodoListFactory(),
	}
}

type service struct {
	todoLists domain.IRepository
	factory   domain.IFactory
}

// GetUserLists returns user's TODO lists slice.
func (s *service) GetUserLists(ctx context.Context, userId uuid.UUID) ([]*domain.TodoList, error) {
	lists, err := s.todoLists.GetByUserID(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("request from repository: %w", err)
	}
	return lists, nil
}

// Create creates new TODO list.
func (s *service) Create(ctx context.Context, user domain.User) (*domain.TodoList, error) {
	newList := s.factory.CreateEmpty(user)
	return newList, s.todoLists.Put(ctx, newList)
}

// PutItem puts new item into TODO list.
func (s *service) PutItem(ctx context.Context, listId uuid.UUID, item *domain.TodoItem) error {
	list, err := s.todoLists.GetByID(ctx, listId)
	if err != nil {
		return fmt.Errorf("get TODO list by ID: %w", err)
	}

	list.AddItem(*item)
	if err = s.todoLists.Update(ctx, list); err != nil {
		return fmt.Errorf("save TODO list updates: %w", err)
	}
	return nil
}
