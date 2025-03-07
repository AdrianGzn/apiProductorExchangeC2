package application

import (
	"productor/src/orders/domain"
	"log"
)

type CreateOrderUseCase struct {
	rabbitqmRepository domain.IOrderRabbitqm
	mysqlRepository domain.IOrderMysq
}

func NewCreateOrderUseCase(rabbitqmRepository domain.IOrderRabbitqm, mysqlRepository domain.IOrderMysq) *CreateOrderUseCase {
	return &CreateOrderUseCase{rabbitqmRepository: rabbitqmRepository, mysqlRepository: mysqlRepository}
}

func (usecase *CreateOrderUseCase) SetOrder(mysqlRepository domain.IOrderMysq, rabbitqmRepository domain.IOrderRabbitqm) {
	usecase.mysqlRepository = mysqlRepository
	usecase.rabbitqmRepository = rabbitqmRepository
}

func (usecase *CreateOrderUseCase) Run(order *domain.Order) error {
	err := usecase.mysqlRepository.Save(order)
	errSendMessage := usecase.rabbitqmRepository.Save(order)

	if err != nil || errSendMessage != nil {
		log.Panicf("error to send message %s", err); 
		return err
	}
	
	return errSendMessage; 
}
