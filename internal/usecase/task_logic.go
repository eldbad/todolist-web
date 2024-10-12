package usecase

import (
	"time"

	"github.com/eldbad/todolist-web/internal/entity"
)

type taskRepository interface {
	FindAll() ([]entity.Task, error)
	Find(*entity.Task) (*entity.Task, error)
	Create(*entity.Task) error
	Update(*entity.Task) error
	Delete(*entity.Task) error
}

type TaskUsecase struct {
	tr taskRepository
}

func NewTaskUsecase(tr taskRepository) *TaskUsecase {
	return &TaskUsecase{tr: tr}
}

func (tu TaskUsecase) GetAllTasks() ([]entity.Task, error) {
	return tu.tr.FindAll()
}

func (tu TaskUsecase) GetTaskById(id int) (*entity.Task, error) {
	task := entity.Task{Id: id}

	return tu.tr.Find(&task)
}

func (tu TaskUsecase) CreateTask(
	name string,
	description string,
	creationDate time.Time,
	status bool,
) error {
	task := entity.Task{
		Name:         name,
		Description:  description,
		CreationDate: creationDate,
		Status:       status,
	}

	return tu.tr.Create(&task)
}

func (tu TaskUsecase) UpdateTaskById(id int) error {
	task := entity.Task{Id: id}

	return tu.tr.Update(&task)
}

func (tu TaskUsecase) DeleteTaskById(id int) error {
	task := entity.Task{Id: id}

	return tu.tr.Delete(&task)
}
