package infrastructure

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "time"

    amqp "github.com/rabbitmq/amqp091-go"
    "productor/src/orders/domain"
)

type MyExchangeLogs struct {
    ch *amqp.Channel
}

func NewRabbitRepository(ch *amqp.Channel) *MyExchangeLogs {
    if err := ch.ExchangeDeclare(
        "logs",   // Nombre del exchange
        "fanout", // Tipo del exchange
        true,     // Durable
        false,    // Auto-deleted
        false,    // Internal
        false,    // No-wait
        nil,      // Argumentos
    ); err != nil {
        log.Fatalf("Error al declarar el exchange: %v", err)
    }

    return &MyExchangeLogs{ch: ch}
}

func (ch *MyExchangeLogs) Save(order *domain.Order) error {
    body, err := json.Marshal(order)
    if err != nil {
        return fmt.Errorf("error al serializar el pedido: %v", err)
    }

    log.Printf("Enviando mensaje: %s", body)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := ch.ch.PublishWithContext(ctx,
        "logs",  // Exchange
        "",      // Routing key
        false,   // Mandatory
        false,   // Immediate
        amqp.Publishing{
            ContentType: "application/json", // Tipo de contenido
            Body:        body,               // Cuerpo del mensaje
        }); err != nil {
        return fmt.Errorf("error al enviar el mensaje a RabbitMQ: %v", err)
    }

    log.Printf(" [x] Enviado: %s", body)
    return nil
}