package main

import (
	"os"

	"github.com/eldbad/todolist-web/internal/delivery/postgresql"
	"github.com/eldbad/todolist-web/internal/handlers"
	"github.com/eldbad/todolist-web/internal/logging"
	"github.com/eldbad/todolist-web/internal/repository"
	"github.com/eldbad/todolist-web/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {

	}
	r := gin.Default()
	dsn, ok := os.LookupEnv("DSN")
	if ok {
		// TODO: log it
	}

	logging.NewLogger()
	db, err := postgresql.InitDB(dsn)
	if err != nil {
		// TODO: log it
	}

	tr := repository.NewTaskRepository(db)
	tu := usecase.NewTaskUsecase(tr)
	th := handlers.NewTaskHandler(tu)
	th.Register(&r.RouterGroup)

	r.Run()

	// log.Error().Stack().Err(err).Msg("")
}
