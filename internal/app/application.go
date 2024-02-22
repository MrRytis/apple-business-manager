package app

import (
	"github.com/MrRytis/apple-business-manager/internal/client"
	"github.com/MrRytis/apple-business-manager/internal/queue"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Application struct {
	DB     *sqlx.DB
	Rabbit *queue.Rabbit
	Logger echo.Logger
	Client client.Client
}

func NewApplication(db *sqlx.DB, r *queue.Rabbit, logger echo.Logger, client client.Client) *Application {
	return &Application{
		DB:     db,
		Rabbit: r,
		Logger: logger,
		Client: client,
	}
}

func (a *Application) AddConsumer(name string) error {
	for _, availableConsumer := range a.Rabbit.AllPossibleConsumers {
		if availableConsumer.GetQueueName() == name {
			con, err := availableConsumer.NewConsumer(a.Rabbit.Conn, a.DB, a.Logger, a.Client)
			if err != nil {
				return errors.Wrap(err, "failed to create new consumer")
			}
			a.Rabbit.Consumers = append(a.Rabbit.Consumers, con)
			return nil
		}

	}

	return errors.New("consumer not found")
}

func (a *Application) RemoveConsumer(name string) error {
	isRemoved := false
	for _, con := range a.Rabbit.Consumers {
		if con.GetQueueName() == name {
			con.Close()
			isRemoved = true
		}
	}

	if !isRemoved {
		return errors.New("consumer not found")
	}

	return nil
}
