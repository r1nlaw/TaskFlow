package repository

type Authorized interface {
}
type TaskList interface {
}
type TaskItem interface {
}

type Repository struct {
	Authorized
	TaskList
	TaskItem
}

func NewRepository() *Repository {
	return &Repository{}
}
