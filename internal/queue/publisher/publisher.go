package publisher

import (
	"encoding/json"
	"github.com/MrRytis/apple-business-manager/internal/queue/message"
	"github.com/wagslane/go-rabbitmq"
	"os"
)

type Publisher struct {
	publisher *rabbitmq.Publisher
}

func (p *Publisher) Close() {
	p.publisher.Close()
}

func New(conn *rabbitmq.Conn) (*Publisher, error) {
	var opts func(*rabbitmq.PublisherOptions)

	if os.Getenv("APP") == "dev" {
		opts = rabbitmq.WithPublisherOptionsLogging
	}

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		opts,
	)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		publisher: publisher,
	}, nil
}

func (p *Publisher) Publish(message message.Message) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		b,
		[]string{message.GetRoutingKey()},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("events"),
	)
	if err != nil {
		return err
	}

	return nil
}
