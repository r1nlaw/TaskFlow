package service

import "TaskFlow/package/repository"

type Authorized interface {
}
type TaskList interface {
}
type TaskItem interface {
}

type Service struct {
	Authorized
	TaskList
	TaskItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
