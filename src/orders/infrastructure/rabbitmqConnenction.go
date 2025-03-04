package infrastructure

import (
	"context"
	"encoding/json"
	"log"
	"time"
    "os"
    "fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"productor/src/orders/domain"
    "github.com/joho/godotenv"
)

type RabbitMQRepository struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

//Para el .env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
}

func NewRabbitMQRepository() (*RabbitMQRepository, error) {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, host, port)

        conn, err := amqp.Dial(rabbitMQURL)
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