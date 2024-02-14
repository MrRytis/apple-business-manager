package main

import (
	"github.com/MrRytis/apple-business-manager/internal/exception"
	"github.com/MrRytis/apple-business-manager/internal/handler"
	"github.com/MrRytis/apple-business-manager/internal/queue"
	"github.com/MrRytis/apple-business-manager/internal/router"
	"github.com/MrRytis/apple-business-manager/internal/storage"
	"github.com/MrRytis/apple-business-manager/internal/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	e.Validator = validator.New()
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

	r, err := queue.New(conn)
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := handler.New(db, r)
	router.RegisterRoutes(e, h)

	e.Logger.Fatal(e.Start(":1323"))
}