package main

import (
	"github.com/MrRytis/apple-business-manager/internal/client"
	"github.com/MrRytis/apple-business-manager/internal/exception"
	"github.com/MrRytis/apple-business-manager/internal/handler"
	"github.com/MrRytis/apple-business-manager/internal/queue"
	"github.com/MrRytis/apple-business-manager/internal/queue/consumer"
	"github.com/MrRytis/apple-business-manager/internal/router"
	"github.com/MrRytis/apple-business-manager/internal/storage"
	"github.com/MrRytis/apple-business-manager/internal/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	e.Validator = validator.NewCustomValidator()
	e.HTTPErrorHandler = exception.CustomHTTPErrorHandler

	db, err := storage.New()
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	conn, err := queue.NewConn()
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer conn.Close()

	r, err := queue.NewRabbit(conn)
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := handler.NewHandler(db, r)
	router.RegisterRoutes(e, h)

	abmClient := client.NewAbmClient(os.Getenv("ABM_URL"), os.Getenv("ABM_SECRET"))

	consumers, err := consumer.StartConsumers(conn, db, e.Logger, abmClient)
	if err != nil {
		e.Logger.Fatal(err)
	}

	defer func() {
		for _, c := range consumers {
			c.Close()
		}
	}()

	e.Logger.Fatal(e.Start(":1323"))
}
