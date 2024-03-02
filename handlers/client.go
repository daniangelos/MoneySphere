package handlers

import (
	"fmt"
	"net/http"

	"github.com/MoneySphere/controllers"
	"github.com/MoneySphere/dtos"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	clientController controllers.IClientController
}

func NewClientHandler(controllerContainer controllers.ControllerContainer) *ClientHandler {
	return &ClientHandler{
		clientController: controllerContainer.ClientController,
	}
}

func (ch *ClientHandler) CreateClientTransaction(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var body dtos.TransactionCreateRequest

	if err := c.BindJSON(&body); err != nil {
		return
	}

	fmt.Println(body)

	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}
