package repository

import (
	"context"
	"sync"

	domain "github.com/exepirit/todo-app/internal/domain/todolist"

	"github.com/google/uuid"
)

// NewMemoryTodoList creates new repository with in-memory storage.
func NewMemoryTodoList() domain.IRepository {
	return &memoryTodoListRepo{
		store: make(map[uuid.UUID]domain.TodoList),
	}
}

type memoryTodoListRepo struct {
	store map[uuid.UUID]domain.TodoList
	lock  sync.RWMutex
}

func (repo *memoryTodoListRepo) GetByUserID(_ context.Context, userId uuid.UUID) ([]*domain.TodoList, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()

	var result []*domain.TodoList
	for _, todolist := range repo.store {
		if todolist.User().ID == userId {
			result = append(result, &todolist)
		}
	}

	return result, nil
}

func (repo *memoryTodoListRepo) GetByID(_ context.Context, id uuid.UUID) (*domain.TodoList, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()

	todolist, ok := repo.store[id]
	if !ok {
		return nil, ErrNotFound
	}
	return &todolist, nil
}

func (repo *memoryTodoListRepo) Put(_ context.Context, todolist *domain.TodoList) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	_, exists := repo.store[todolist.ID()]
	if exists {
		return ErrExists
	}

	repo.store[todolist.ID()] = *todolist
	return nil
}

func (repo *memoryTodoListRepo) Update(_ context.Context, todolist *domain.TodoList) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	_, exists := repo.store[todolist.ID()]
	if !exists {
		return ErrNotFound
	}

	repo.store[todolist.ID()] = *todolist
	return nil
}
