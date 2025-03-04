package infrastructure

import (
	"productor/src/orders/application"
	"productor/src/orders/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
	useCase *application.CreateOrderUseCase
}

func NewCreateOrderController(useCase *application.CreateOrderUseCase) *CreateOrderController {
	return &CreateOrderController{useCase: useCase}
}

func (controller *CreateOrderController) Execute(c *gin.Context) {
	var order domain.Order
	
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos"})
		return
	}

	if err := controller.useCase.Run(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orden creada y enviada a RabbitMQ"})
}
