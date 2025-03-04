package infrastructure

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"productor/src/orders/domain"
)

type RabbitMQRepository struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQRepository() (*RabbitMQRepository, error) {
	conn, err := amqp.Dial("amqp://adri:1234@174.129.127.24:5672")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"logs",   // nombre
		"fanout", // tipo
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQRepository{conn: conn, ch: ch}, nil
}

//para mandarlo a rabbit
func (repo *RabbitMQRepository) Save(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = repo.ch.PublishWithContext(ctx,
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json", // en vez de texto va JSON
			Body:        body,
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s", body)
	return nil
}

func (repo *RabbitMQRepository) Close() {
	if repo.ch != nil {
		repo.ch.Close()
	}
	if repo.conn != nil {
		repo.conn.Close()
	}
}


/* No necesario pq se convierte a json
func bodyFrom(args []string) string {
        var s string
        if (len(args) < 2) || os.Args[1] == "" {
                s = "hello"
        } else {
                s = strings.Join(args[1:], " ")
        }
        return s
}
*/