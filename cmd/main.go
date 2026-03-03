package main

import (
	"github.com/VaLTrexx/crud/internal/handler"
	"github.com/VaLTrexx/crud/internal/repository"
	"github.com/VaLTrexx/crud/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	db := repository.Newdb()
	defer db.Close()

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()

	router.GET("/tasks", taskHandler.GetAllTasks)
	router.POST("/tasks", taskHandler.CreateTask)

	router.Run()
}
