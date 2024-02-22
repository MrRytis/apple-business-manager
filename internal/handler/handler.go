package handler

import (
	"github.com/MrRytis/apple-business-manager/internal/app"
)

type Handler struct {
	App *app.Application
}

func NewHandler(app *app.Application) *Handler {
	return &Handler{
		App: app,
	}
}
