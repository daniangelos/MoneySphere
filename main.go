package main

import (
	"fmt"
	"net/http"

	"github.com/MoneySphere/config"
	"github.com/MoneySphere/controllers"
	"github.com/MoneySphere/handlers"
	"github.com/MoneySphere/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := config.NewConfig()

	if err != nil {
		fmt.Println("config error", err)
		return
	}

	db, err := initDB(config.DB)

	if err != nil {
		fmt.Println("db error", err)
		return
	}

	repositoryContainer := repositories.NewRepositoryContainer(db)
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

func initDB(dbConfig config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, err
}
