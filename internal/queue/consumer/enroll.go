package consumer

import (
	"database/sql"
	"encoding/json"
	"github.com/MrRytis/apple-business-manager/internal/client"
	"github.com/MrRytis/apple-business-manager/internal/factory"
	"github.com/MrRytis/apple-business-manager/internal/queue/message"
	"github.com/MrRytis/apple-business-manager/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/wagslane/go-rabbitmq"
	"time"
)

type EnrollConsumer struct {
	Consumer *rabbitmq.Consumer
}

const (
	queue = `enroll.consume`
)

func NewEnrollConsumer(conn *rabbitmq.Conn, db *sqlx.DB, logger echo.Logger, client *client.AbmClient) (RabbitConsumer, error) {
	consumer, err := newConsumer(
		conn,
		message.EnrollRoutingKey,
		queue,
		NewWorker(db, logger, client),
	)

	if err != nil {
		return nil, err
	}

	return &EnrollConsumer{
		Consumer: consumer,
	}, nil
}

func (c *EnrollConsumer) GetQueueName() string {
	return queue
}

func (c *EnrollConsumer) Close() {
	c.Consumer.Close()
}

func (w *Worker) handler(d rabbitmq.Delivery) rabbitmq.Action {
	var enroll message.EnrollMessage

	err := json.Unmarshal(d.Body, &enroll)
	if err != nil {
		w.Logger.Error(err)

		return rabbitmq.NackDiscard
	}

	//TODO: Update query
	trans, err := storage.GetTransactionById(w.Db, enroll.TransactionId)
	if err != nil {
		w.Logger.Error(err)

		return rabbitmq.NackDiscard
	}

	appleReq := factory.CreateEnrollRequest(trans)
	appleRes, err := w.Client.BulkEnroll(appleReq)

	//TODO: handle error
	if err != nil {
		w.Logger.Error(err)

		return rabbitmq.NackRequeue

	}

	trans.Status = storage.TransactionProcessing
	trans.AppleTransactionId = sql.NullString{
		String: appleRes.AppleTransactionId,
		Valid:  true,
	}
	trans.UpdatedAt = time.Now()

	err = storage.UpdateTransactionStatusAndAppleId(w.Db, trans)

	return rabbitmq.Ack
}
