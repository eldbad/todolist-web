package postgresql

import (
	"context"
	"database/sql"

	"github.com/eldbad/todolist-web/internal/entity"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func init() {
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	sqld := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB = bun.NewDB(sqld, pgdialect.New())
	ctx := context.Background()
	DB.NewCreateTable().Model((*entity.Task)(nil)).Exec(ctx)
}
