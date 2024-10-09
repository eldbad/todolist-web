package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/eldbad/todolist-web/internal/entity"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func InitDB(dsn string) (*bun.DB, error) {
	sqld := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB := bun.NewDB(sqld, pgdialect.New())
	ctx := context.Background()
	DB.NewCreateTable().Model((*entity.Task)(nil)).Exec(ctx)

	return DB, errors.New("oops") // TODO: make appropriate custom error
}
