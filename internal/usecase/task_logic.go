package usecase

import (
	"github.com/eldbad/todolist-web/internal/entity"
	"github.com/eldbad/todolist-web/internal/repository"
)

type TaskUsecase struct {
	tr *repository.TaskRepository
}

func NewTaskUsecase(tr *repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{tr: tr}
}

func (tu TaskUsecase) GetAllTasks() ([]entity.Task, error) {
	return tu.tr.FindAll()
}
