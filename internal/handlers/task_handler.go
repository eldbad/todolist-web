package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eldbad/todolist-web/internal/entity"
)

type taskHandler struct {
	taskUsecase taskUsecase
}

type taskUsecase interface {
	GetAllTasks() ([]entity.Task, error)
}

func NewTaskHandler(tl taskUsecase) *taskHandler {
	return &taskHandler{
		taskUsecase: tl,
	}
}

func (th taskHandler) Register(r *gin.RouterGroup) {
	r.GET("/get_task", th.GetTask) // TODO: Make it the same with other methods
}

func (th *taskHandler) GetTask(c *gin.Context) {
	tasks, err := th.taskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (th *taskHandler) GetTasks(c *gin.Context) {
	panic("not implemented")
}

func (th *taskHandler) CreateTask(c *gin.Context) {
	panic("not implemented")
}

func (th *taskHandler) UpdateTask(c *gin.Context) {
	panic("not implemented")
}

func (th *taskHandler) DeleteTask(c *gin.Context) {
	panic("not implemented")
}
