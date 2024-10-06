package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eldbad/todolist-web/internal/entity"
)

type TaskHandler struct {
	taskLogic taskUsecase
}

type taskUsecase interface {
	GetAllTasks() ([]entity.Task, error)
}

func NewTaskHandler(tl taskUsecase) *TaskHandler {
	return &TaskHandler{
		taskLogic: tl,
	}
}

func (th TaskHandler) Register(r *gin.RouterGroup) {
	r.GET("/get_task", th.GetTask) // TODO: Make it the same with other methods
}

func (th *TaskHandler) GetTask(c *gin.Context) {
	tasks, err := th.taskLogic.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (th *TaskHandler) GetTasks(c *gin.Context) {
	panic("not implemented")
}

func (th *TaskHandler) CreateTask(c *gin.Context) {
	panic("not implemented")
}

func (th *TaskHandler) UpdateTask(c *gin.Context) {
	panic("not implemented")
}

func (th *TaskHandler) DeleteTask(c *gin.Context) {
	panic("not implemented")
}
