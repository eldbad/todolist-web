package main

import (
	"github.com/eldbad/todolist-web/internal/handlers"
	"github.com/eldbad/todolist-web/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	tl := usecase.TaskUsecase{}
	th := handlers.NewTaskHandler(tl)
	th.Register(&r.RouterGroup)

	r.Run()

	// log.Error().Stack().Err(err).Msg("")
}
