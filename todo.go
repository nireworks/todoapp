package todoapp

type TodoApp struct {
}

func New() *TodoApp {
	return &TodoApp{}
}

func (t *TodoApp) GetTodo(int) (Todo, error) {
	panic("implement me")
}

func (t *TodoApp) GetTodos() ([]Todo, error) {
	panic("implement me")
}

func (t *TodoApp) SaveTodo(Todo) error {
	panic("implement me")
}
