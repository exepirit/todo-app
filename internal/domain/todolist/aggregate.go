package todolist

import "github.com/google/uuid"

// TodoList is a aggregate describes user's TODO list structure.
type TodoList struct {
	id    uuid.UUID
	user  User
	items []TodoItem
}

// ID returns identifier of a aggregate.
func (todolist *TodoList) ID() uuid.UUID {
	if todolist.id == uuid.Nil {
		todolist.id = uuid.New()
	}
	return todolist.id
}

// User returns user entity.
func (todolist *TodoList) User() User {
	return todolist.user
}

// User sets provided user as TODO list owner.
func (todolist *TodoList) SetUser(user User) {
	todolist.user = user
}

// Items returns slice of stored items.
func (todolist *TodoList) Items() []TodoItem {
	return todolist.items
}

// AddItem adds item to TODO list.
func (todolist *TodoList) AddItem(item TodoItem) {
	todolist.items = append(todolist.items, item)
}
