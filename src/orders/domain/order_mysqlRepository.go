package domain

type IOrderMysq interface {
	Save(order *Order) error
}
