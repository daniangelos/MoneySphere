package handlers

import "github.com/MoneySphere/controllers"

type HandlerContainer struct {
	ClientHandler *ClientHandler
}

func NewHandlerContainer(controllerContainer controllers.ControllerContainer) HandlerContainer {
	return HandlerContainer{
		ClientHandler: NewClientHandler(controllerContainer),
	}
}
