package factory

import (
	"github.com/MrRytis/apple-business-manager/internal/client"
	"github.com/MrRytis/apple-business-manager/internal/storage"
	"github.com/google/uuid"
	"time"
)

func CreateEnrollRequest(trans *storage.Transaction) *client.BulkEnrollRequest {
	var devices []*client.Device
	for _, device := range trans.Order.Devices {
		devices = append(devices, &client.Device{
			Imei: device.Imei,
		})
	}

	var deliveries []*client.Delivery
	deliveries = append(deliveries, &client.Delivery{
		DeliveryNumber: uuid.New().String(),
		ShipDate:       trans.Order.ContractDate.UTC().Format(time.RFC3339),
		Devices:        devices,
	})

	var orders []*client.Order

	orders = append(orders, &client.Order{
		OrderNumber: trans.Order.UUID,
		OrderDate:   trans.Order.ContractDate.UTC().Format(time.RFC3339),
		OrderType:   string(trans.Status),
		CustomerId:  trans.Order.CustomerId,
		Deliveries:  deliveries,
	})

	req := &client.BulkEnrollRequest{
		RequestContext: client.RequestContext{
			ShipTo:   "",
			LangCode: "en",
			TimeZone: "Europe/Vilnius",
		},
		TransactionId: trans.UUID,
		DepResellerId: "",
		Orders:        orders,
	}

	return req
}
