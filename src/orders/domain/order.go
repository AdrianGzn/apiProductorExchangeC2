package domain

type Order struct {
	Name  string  		`json:"name"`
	Description string 	`json:"description"`
	Price int32 		`json:"price"`
	UserName string		`json:"userName"`
	UserCellphone string`json:"cellPhone"`
}

func NewOrder(name string, description string, price int32, userName string, userCellphone string) *Order {
	return &Order{Name: name, Description: description, Price: price, UserName: userName, UserCellphone: userCellphone}
}