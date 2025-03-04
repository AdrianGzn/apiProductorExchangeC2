package domain

type IOrder interface {
	Save(order *Order) error
}
