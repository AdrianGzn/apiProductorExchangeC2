package infrastructure

import (
	"productor/src/orders/domain"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *MysqlRepository {
	return &MysqlRepository{db: db}
}

func (repo *MysqlRepository) Save(order *domain.Order) error {

	query := "INSERT INTO orders (name, description, price, userName, userCellphone) VALUES (?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, order.Name, order.Description, order.Price, order.UserName, order.UserCellphone)
	if err != nil {
		return fmt.Errorf("Error al guardar orden: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error al obtener el ID insertado: %v", err)
	}
	order.Id = int32(id)

	return nil
}