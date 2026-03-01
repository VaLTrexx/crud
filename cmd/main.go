package main

import (
	"net/http"

	"github.com/VaLTrexx/crud/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	db := repository.Newdb()
	defer db.Close()

	taskRepo := repository.NewTaskRepository(db)
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "runing",
		})
	})

	router.Run()
}
