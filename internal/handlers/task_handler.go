package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/eldbad/todolist-web/internal/entity"
)

type taskHandler struct {
	taskUsecase taskUsecase
}

type taskUsecase interface {
	GetAllTasks() ([]entity.Task, error)
	GetTaskById(int) (*entity.Task, error)
	CreateTask(string, string, time.Time, bool) error
	UpdateTaskById(int) error
	DeleteTaskById(int) error
}

func NewTaskHandler(tl taskUsecase) *taskHandler {
	return &taskHandler{
		taskUsecase: tl,
	}
}

func (th taskHandler) Register(r *gin.RouterGroup) {
	r.GET("/task", th.GetTasks)
	r.GET("/task/:id", th.GetTask)
	r.POST("/task/:name", th.CreateTask)
	r.PUT("/task/:id", th.UpdateTask)
	r.DELETE("/task/:id", th.DeleteTask)
}

func (th *taskHandler) GetTasks(c *gin.Context) {
	tasks, err := th.taskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (th *taskHandler) GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	task, err := th.taskUsecase.GetTaskById(id)
	if err != nil {
		// TODO: think about error handling in this case
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (th *taskHandler) CreateTask(c *gin.Context) {
	err := th.taskUsecase.CreateTask(c.Param("name"), "", time.Now(), false)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{})
}

func (th *taskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	err = th.taskUsecase.UpdateTaskById(id)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{})
}

func (th *taskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	err = th.taskUsecase.DeleteTaskById(id)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{})
}
