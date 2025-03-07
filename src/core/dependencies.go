package core

import (
	"log"
	"productor/src/orders/application"
	"productor/src/orders/infrastructure"
	"productor/src/core/middlewares"
	"github.com/gin-gonic/gin"
)

func IniciarRutas() {
    mysqlConn, err := GetDBPool()
    if err != nil {
        log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
    }

	rabbitmqCh, err := GetChannel()
	if err != nil {
        log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
    }

    mysqlRepository := infrastructure.NewMysqlRepository(mysqlConn.DB)
	rabbitqmRepository := infrastructure.NewRabbitRepository(rabbitmqCh.ch)

	createOrderUseCase := application.NewCreateOrderUseCase(rabbitqmRepository, mysqlRepository)
	createOrderController := infrastructure.NewCreateOrderController(createOrderUseCase)

	router := gin.Default()
	middleware := middlewares.NewCorsMiddleware()	
	router.Use(middleware)

	router.POST("/order", createOrderController.Execute)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}