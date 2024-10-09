package main

import (
	"os"

	"github.com/eldbad/todolist-web/internal/delivery/postgresql"
	"github.com/eldbad/todolist-web/internal/handlers"
	"github.com/eldbad/todolist-web/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dsn, ok := os.LookupEnv("DSN")
	if ok {
		// TODO: log it
	}

	postgresql.InitDB(dsn)

	tl := usecase.TaskUsecase{}
	th := handlers.NewTaskHandler(tl)
	th.Register(&r.RouterGroup)

	r.Run()

	// log.Error().Stack().Err(err).Msg("")
}
