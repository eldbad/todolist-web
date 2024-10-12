package usecase

import (
	"time"

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
