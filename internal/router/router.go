package router

import (
	"github.com/MrRytis/apple-business-manager/internal/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, h *handler.Handler) {
	e.POST("/api/v1/enroll", h.Enroll)

	e.GET("/api/v1/consumers", h.ListConsumers)
	e.POST("/api/v1/consumer", h.AddConsumer)
	e.DELETE("/api/v1/consumers/:id", h.RemoveConsumer)
}
