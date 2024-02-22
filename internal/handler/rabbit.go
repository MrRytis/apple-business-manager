package handler

import (
	"github.com/MrRytis/apple-business-manager/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) ListConsumers(c echo.Context) error {
	var list map[string]int
	for _, con := range h.App.Rabbit.Consumers {
		value, exists := list[con.GetQueueName()]
		if exists {
			list[con.GetQueueName()] = value + 1
		} else {
			list[con.GetQueueName()] = 1
		}
	}

	var consumers []model.ConsumerResponse
	for name, count := range list {
		consumers = append(consumers, model.ConsumerResponse{
			QueueName: name,
			Count:     count,
		})
	}

	return c.JSON(http.StatusOK, consumers)
}

func (h *Handler) AddConsumer(c echo.Context) error {
	var req model.AddConsumerRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	err := h.App.AddConsumer(req.QueueName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}

func (h *Handler) RemoveConsumer(c echo.Context) error {
	err := h.App.RemoveConsumer(c.Param("name"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "consumer not found")
	}

	return c.JSON(http.StatusNoContent, nil)
}
