package repository

import (
	"context"

	"github.com/eldbad/todolist-web/internal/delivery/postgresql"
	"github.com/eldbad/todolist-web/internal/entity"
	"github.com/uptrace/bun"
)

type TaskRepository struct {
}

func NewTaskRepository(db *bun.DB) *TaskRepository {
	tr := &TaskRepository{}

	return tr
}

func (tr *TaskRepository) FindAll() ([]entity.Task, error) {
	tasks := []entity.Task{}
	ctx := context.Background()

	err := postgresql.DB.NewSelect().Model(&tasks).Scan(ctx)
	if err != nil {
		// TODO: log an error
	}

	return tasks, err // TODO: return custom error or use logging
}

func (tr *TaskRepository) FindById(id int) (entity.Task, error) {
	task := entity.Task{}
	ctx := context.Background()
	err := postgresql.DB.NewSelect().Model(&task).Where("id = ?", id).Scan(ctx)

	return task, err
}

func (tr *TaskRepository) Create(task *entity.Task) error {
	ctx := context.Background()
	_, err := postgresql.DB.NewInsert().Model(task).Exec(ctx)

	return err
}

func (tr *TaskRepository) Update(id int) error {
	task := entity.Task{Id: id}
	ctx := context.Background()

	_, err := postgresql.DB.NewUpdate().Model(task).WherePK().Exec(ctx)

	return err
}

func (tr *TaskRepository) Delete(id int) error {
	task := entity.Task{Id: id}
	ctx := context.Background()

	_, err := postgresql.DB.NewDelete().Model(task).WherePK().Exec(ctx)

	return err
}
