package queue

import (
	"fmt"
	"github.com/MrRytis/apple-business-manager/internal/queue/consumer"
	"github.com/MrRytis/apple-business-manager/internal/queue/publisher"
	"github.com/pkg/errors"
	"github.com/wagslane/go-rabbitmq"
	"os"
)

type Rabbit struct {
	Conn                 *rabbitmq.Conn
	Publisher            *publisher.Publisher
	Consumers            []consumer.RabbitConsumer
	AllPossibleConsumers []consumer.RabbitConsumer
}

func NewRabbit() (*Rabbit, error) {
	conn, err := newConn()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new connection")
	}

	pub, err := publisher.New(conn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new publisher")
	}

	return &Rabbit{
		Conn:                 conn,
		Publisher:            pub,
		AllPossibleConsumers: getPossibleConsumers(),
	}, nil
}

func (r *Rabbit) Close() {
	r.Publisher.Close()

	for _, con := range r.Consumers {
		con.Close()
	}

	_ = r.Conn.Close()
}

func newConn() (*rabbitmq.Conn, error) {
	var opts func(*rabbitmq.ConnectionOptions)

	if os.Getenv("APP") == "dev" {
		opts = rabbitmq.WithConnectionOptionsLogging
	}

	conn, err := rabbitmq.NewConn(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%s",
			os.Getenv("RABBIT_USER"),
			os.Getenv("RABBIT_PASSWORD"),
			os.Getenv("RABBIT_HOST"),
			os.Getenv("RABBIT_PORT"),
		),
		opts,
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func getPossibleConsumers() []consumer.RabbitConsumer {
	return []consumer.RabbitConsumer{
		&consumer.EnrollConsumer{},
	}
}
