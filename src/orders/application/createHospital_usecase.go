package application

import "productor/src/orders/domain"

type CreateOrderUseCase struct {
	repo domain.IOrder
}

func NewCreateOrderUseCase(repo domain.IOrder) *CreateOrderUseCase {
	return &CreateOrderUseCase{repo: repo}
}

func (usecase *CreateOrderUseCase) Run(order *domain.Order) error {
	return usecase.repo.Save(order)  // guarda el pedido y lo envía a rab
}
