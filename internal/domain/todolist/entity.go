package todolist

import "github.com/google/uuid"

// User structure represents basic user's data.
type User struct {
	ID   uuid.UUID
	Name string
}

// TodoItem represents single item of TODO list.
type TodoItem struct {
	Text string
}
