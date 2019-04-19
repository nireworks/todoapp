package model

// Todo is the underlying structure that is bein handled by the TodoService
type Todo struct {
	Id        int
	Title     string
	Completed bool
}
