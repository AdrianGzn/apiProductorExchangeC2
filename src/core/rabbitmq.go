package core

import (
    "fmt"
    "log"
    "os"

    amqp "github.com/rabbitmq/amqp091-go"
    "github.com/joho/godotenv"
)

type RabbitMQRepository struct {
    conn *amqp.Connection
    ch   *amqp.Channel
}

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }
}

func GetChannel() (*RabbitMQRepository, error) {
    rabbitMQUser := os.Getenv("RABBITMQ_USER")
    rabbitMQPassword := os.Getenv("RABBITMQ_PASSWORD")
    rabbitMQHost := os.Getenv("RABBITMQ_HOST")
    rabbitMQPort := os.Getenv("RABBITMQ_PORT")

    log.Printf("Conectando con las siguientes credenciales: %s:%s@%s:%s", rabbitMQUser, rabbitMQPassword, rabbitMQHost, rabbitMQPort)

    rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%s", rabbitMQUser, rabbitMQPassword, rabbitMQHost, rabbitMQPort)

    conn, err := amqp.Dial(rabbitMQURL)
    if err != nil {
        return nil, fmt.Errorf("error al conectar con RabbitMQ: %v", err)
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, fmt.Errorf("error al abrir un canal: %v", err)
    }

    return &RabbitMQRepository{ch: ch}, nil
}

func (repo *RabbitMQRepository) Close() {
    if repo.ch != nil {
        if err := repo.ch.Close(); err != nil {
            log.Printf("Error al cerrar el canal: %v", err)
        } else {
            log.Println("Canal cerrado")
        }
    }

    if repo.conn != nil {
        if err := repo.conn.Close(); err != nil {
            log.Printf("Error al cerrar la conexión: %v", err)
        } else {
            log.Println("Conexión cerrada")
        }
    }
}