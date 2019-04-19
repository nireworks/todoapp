package todoapp

type TodoService interface {
	GetTodo(int) (Todo, error)
	GetTodos() ([]Todo, error)
	SaveTodo(Todo) error
}

// Todo is the underlying structure that is bein handled by the TodoService
type Todo struct {
	Id        int
	Title     string
	Completed bool
}
