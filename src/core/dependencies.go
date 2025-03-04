package core

import (
	"log"

	"productor/src/orders/application"
	"productor/src/orders/infrastructure"
	"productor/src/core/middlewares"

	"github.com/gin-gonic/gin"
)

func IniciarRutas() {
	repo, err := infrastructure.NewRabbitMQRepository()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}

	createOrderUseCase := application.NewCreateOrderUseCase(repo)
	createOrderController := infrastructure.NewCreateOrderController(createOrderUseCase);

	router := gin.Default()
	middleware := middlewares.NewCorsMiddleware()	
	router.Use(middleware)

	router.POST("/order", createOrderController.Execute)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}