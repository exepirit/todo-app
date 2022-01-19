package todolist

// IFactory encapsulates aggregate creation.
type IFactory interface {
	// CreateEmpty creates new TodoList aggregate instance.
	CreateEmpty(owner User) *TodoList
}

// NewTodoListFactory creates new factory.
func NewTodoListFactory() IFactory {
	return &factory{}
}

type factory struct{}

func (factory) CreateEmpty(owner User) *TodoList {
	return &TodoList{
		user: owner,
	}
}
