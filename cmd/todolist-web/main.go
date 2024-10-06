package main

import (
	"github.com/eldbad/todolist-web/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	th := handlers.TaskHandler{}

	r.GET("/get_tasks", th.GetTask)
	r.Run()

	log.Error().Stack().Err(err).Msg("")
}
