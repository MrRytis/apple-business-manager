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

	tran := factory.CreatePendingTransaction()
	order := factory.CreateEnrollOrder(req.CustomerId, req.ContractDate, req.Number)
	delivery := factory.CreateDelivery(req.DeliveryNumber)
	devices := factory.CreateDevices(req.Devices)

	t, err := h.DB.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to create transaction")
	}

	tran, err = storage.SaveTransactionTx(t, tran)
	if err != nil {
		_ = t.Rollback()
		return errors.Wrap(err, "failed to save transaction")
	}

	order.TransactionId = tran.Id
	order, err = storage.SaveOrderTx(t, order)
	if err != nil {
		_ = t.Rollback()
		return errors.Wrap(err, "failed to save order")
	}

	delivery.OrderId = order.Id
	delivery, err = storage.SaveDeliveryTx(t, delivery)
	if err != nil {
		_ = t.Rollback()
		return errors.Wrap(err, "failed to save delivery")
	}

	for _, device := range devices {
		device.DeliveryId = delivery.Id
	}

	devices, err = storage.SaveDevicesTx(t, devices)
	if err != nil {
		_ = t.Rollback()
		return errors.Wrap(err, "failed to save devices")
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

	return c.JSON(http.StatusCreated, model.EnrollResponse{Number: tran.Number})
}
