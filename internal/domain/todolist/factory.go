package todolist

// IFactory incapsulates aggregate creation.
type ITodoListFactory interface {
	// CreateEmpty creates new TodoList aggregate instance.
	CreateEmpty(owner User) *TodoList
}

// NewTodoListFactory creates new factory.
func NewTodoListFactory() ITodoListFactory {
	return &factory{}
}

type factory struct{}

func (factory) CreateEmpty(owner User) *TodoList {
	return &TodoList{
		user: owner,
	}
}
