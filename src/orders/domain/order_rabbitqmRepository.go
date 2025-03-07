package domain

type IOrderRabbitqm interface {
	Save(order *Order) error
}
