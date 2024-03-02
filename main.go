package main

import (
	"net/http"

	"github.com/MoneySphere/controllers"
	"github.com/MoneySphere/handlers"
	"github.com/MoneySphere/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	repositoryContainer := repositories.NewRepositoryContainer()
	controllerContainer := controllers.NewControllerContainer(repositoryContainer)
	handlerContainer := handlers.NewHandlerContainer(controllerContainer)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	clientHandler := handlerContainer.ClientHandler

	router.POST("/clientes/:id/transacoes", clientHandler.CreateClientTransaction)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
