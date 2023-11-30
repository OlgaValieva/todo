package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"todo/pkg/service"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}
