package repository

import (
	"context"

	"github.com/eldbad/todolist-web/internal/entity"
	"github.com/uptrace/bun"
)

type TaskRepository struct {
	db *bun.DB
}

func NewTaskRepository(db *bun.DB) *TaskRepository {
	tr := &TaskRepository{db: db}

	return tr
}

func (tr *TaskRepository) FindAll() ([]entity.Task, error) {
	tasks := []entity.Task{}
	ctx := context.Background()

	err := tr.db.NewSelect().Model(&tasks).Scan(ctx)

	return tasks, err
}

func (tr *TaskRepository) Find(task *entity.Task) (*entity.Task, error) {
	ctx := context.Background()
	err := tr.db.NewSelect().Model(&task).Where("id = ?", task.Id).Scan(ctx)

	return task, err
}

func (tr *TaskRepository) Create(task *entity.Task) error {
	ctx := context.Background()
	_, err := tr.db.NewInsert().Model(task).Exec(ctx)

	return err
}

func (tr *TaskRepository) Update(task *entity.Task) error {
	ctx := context.Background()

	_, err := tr.db.NewUpdate().Model(task).WherePK().Exec(ctx)

	return err
}

func (tr *TaskRepository) Delete(task *entity.Task) error {
	ctx := context.Background()

	_, err := tr.db.NewDelete().Model(task).WherePK().Exec(ctx)

	return err
}
