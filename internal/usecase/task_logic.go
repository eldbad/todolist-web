package usecase

import (
	"github.com/eldbad/todolist-web/internal/entity"
)

type TaskUsecase struct{}

func (tl TaskUsecase) GetAllTasks() ([]entity.Task, error) {
	panic("not implemented")
}
