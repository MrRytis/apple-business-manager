package router

import (
	"github.com/MrRytis/apple-business-manager/internal/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, h *handler.Handler) {
	e.POST("/api/v1/enroll", h.Enroll)
}
