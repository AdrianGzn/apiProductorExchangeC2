package core

import (
	"log"
	"productor/src/orders/application"
	"productor/src/orders/infrastructure"
	"productor/src/core/middlewares"
	"github.com/gin-gonic/gin"
	"os"
)

func IniciarRutas() {
	mysqlConn, err := GetDBPool()
	if err != nil {
		log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
	}

	rabbitmqCh, err := GetChannel()
	if err != nil {
		log.Fatalf("Error al obtener la conexión a RabbitMQ: %v", err)
	}

	mysqlRepository := infrastructure.NewMysqlRepository(mysqlConn.DB)
	rabbitqmRepository := infrastructure.NewRabbitRepository(rabbitmqCh.ch)

	createOrderUseCase := application.NewCreateOrderUseCase(rabbitqmRepository, mysqlRepository)
	createOrderController := infrastructure.NewCreateOrderController(createOrderUseCase)

	router := gin.Default()
	router.Use(middlewares.NewCorsMiddleware())
	router.POST("/order", createOrderController.Execute)
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
	
}