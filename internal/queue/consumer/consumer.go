package consumer

import (
	"github.com/MrRytis/apple-business-manager/internal/client"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wagslane/go-rabbitmq"
)

type Worker struct {
	Db     *sqlx.DB
	Logger echo.Logger
	Client *client.AbmClient
}

type RabbitConsumer interface {
	GetQueueName() string
	Close()
}

type WorkerInterface interface {
	handler(d rabbitmq.Delivery) rabbitmq.Action
}

func NewWorker(db *sqlx.DB, logger echo.Logger, c *client.AbmClient) *Worker {
	return &Worker{
		Db:     db,
		Logger: logger,
		Client: c,
	}
}

func StartConsumers(conn *rabbitmq.Conn, db *sqlx.DB, logger echo.Logger, client *client.AbmClient) ([]RabbitConsumer, error) {
	var consumers []RabbitConsumer

	con, err := NewEnrollConsumer(conn, db, logger, client)
	if err != nil {
		return consumers, errors.Wrap(err, "failed to start enroll consumer")
	}

	consumers = append(consumers, con)

	return consumers, nil
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
