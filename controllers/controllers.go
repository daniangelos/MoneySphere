package controllers

import "github.com/MoneySphere/repositories"

type ControllerContainer struct {
	ClientController IClientController
}

func NewControllerContainer(repositoryContainer repositories.RepositoryContainer) ControllerContainer {
	return ControllerContainer{
		ClientController: NewClientController(repositoryContainer),
	}
}
