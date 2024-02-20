package queue

import (
	"fmt"
	"github.com/MrRytis/apple-business-manager/internal/queue/publisher"
	"github.com/pkg/errors"
	"github.com/wagslane/go-rabbitmq"
	"os"
)

type Rabbit struct {
	Conn *rabbitmq.Conn
	Pub  *publisher.Publisher
}

func NewRabbit(conn *rabbitmq.Conn) (*Rabbit, error) {
	pub, err := publisher.New(conn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new publisher")
	}

	return &Rabbit{
		Conn: conn,
		Pub:  pub,
	}, nil
}

func NewConn() (*rabbitmq.Conn, error) {
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
