package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MoneySphere/constants"
	"github.com/MoneySphere/controllers"
	"github.com/MoneySphere/dtos"
	"github.com/MoneySphere/validation"
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
	idText, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		fmt.Println("id not a number")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var body dtos.TransactionCreateRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validation.Validate(body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	transaction, err := ch.clientController.CreateClientTransaction(id, body)

	if err != nil {
		httpStatusCode := http.StatusInternalServerError

		if errors.Is(err, constants.ErrNotEnoughBalance) {
			httpStatusCode = http.StatusUnprocessableEntity
		} else if errors.Is(err, constants.ErrClientNotFound) {
			httpStatusCode = http.StatusNotFound
		}

		fmt.Println("client transaction creation error")
		c.JSON(httpStatusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func (ch *ClientHandler) GetClientTransactions(c *gin.Context) {
	idText, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		fmt.Println("id not a number")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := ch.clientController.GetClientTransactions(id)
	if err != nil {
		httpStatusCode := http.StatusInternalServerError

		if errors.Is(err, constants.ErrClientNotFound) {
			httpStatusCode = http.StatusNotFound
		}

		fmt.Println("failed to get client transactions")
		c.JSON(httpStatusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
