package handler

import (
	"github.com/MrRytis/apple-business-manager/internal/factory"
	"github.com/MrRytis/apple-business-manager/internal/model"
	"github.com/MrRytis/apple-business-manager/internal/queue/message"
	"github.com/MrRytis/apple-business-manager/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) Enroll(c echo.Context) error {
	var req model.EnrollRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	tran := factory.CreatePendingTransaction(req)

	t, err := h.DB.Beginx()
	err = storage.SaveEnrollmentTx(t, tran)
	if err != nil {
		_ = t.Rollback()
		return errors.Wrap(err, "failed to save enrollment")
	}

	m := message.EnrollMessage{
		RoutingKey:    message.EnrollRoutingKey,
		TransactionId: tran.Id,
	}

	err = h.Rabbit.Pub.Publish(&m)
	if err != nil {
		_ = t.Rollback()
		return errors.Wrap(err, "failed to publish enroll message")
	}

	err = t.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return c.JSON(http.StatusCreated, model.EnrollResponse{Number: tran.UUID})
}
