package consumer

import (
	"github.com/MrRytis/apple-business-manager/internal/client"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/wagslane/go-rabbitmq"
)

type RabbitConsumer interface {
	GetQueueName() string
	Close()
	NewConsumer(conn *rabbitmq.Conn, db *sqlx.DB, logger echo.Logger, client client.Client) (RabbitConsumer, error)
}

type WorkerInterface interface {
	handler(d rabbitmq.Delivery) rabbitmq.Action
}

func newConsumer(conn *rabbitmq.Conn, routingKey string, queueName string, worker WorkerInterface) (*rabbitmq.Consumer, error) {
	consumer, err := rabbitmq.NewConsumer(
		conn,
		worker.handler,
		queueName,
		rabbitmq.WithConsumerOptionsRoutingKey(routingKey),
		rabbitmq.WithConsumerOptionsExchangeName("events"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
