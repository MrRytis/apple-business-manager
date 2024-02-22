package main

import (
	"github.com/MrRytis/apple-business-manager/internal/app"
	"github.com/MrRytis/apple-business-manager/internal/client"
	"github.com/MrRytis/apple-business-manager/internal/exception"
	"github.com/MrRytis/apple-business-manager/internal/handler"
	"github.com/MrRytis/apple-business-manager/internal/queue"
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

	db, err := storage.New()
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			e.Logger.Fatal(err)
		}
	}()

	rabbit, err := queue.NewRabbit()
	if err != nil {
		e.Logger.Fatal(err)
	}

	defer func() {
		rabbit.Close()
	}()

	application := app.NewApplication(
		db,
		rabbit,
		e.Logger,
		client.NewAbmClient(
			os.Getenv("ABM_URL"),
			os.Getenv("ABM_SECRET"),
		),
	)

	e.Validator = validator.NewCustomValidator()
	e.HTTPErrorHandler = exception.CustomHTTPErrorHandler

	h := handler.NewHandler(application)
	router.RegisterRoutes(e, h)

	e.Logger.Fatal(e.Start(":1323"))
}
