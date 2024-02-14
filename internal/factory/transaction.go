package factory

import (
	"database/sql"
	"github.com/MrRytis/apple-business-manager/internal/model"
	"github.com/MrRytis/apple-business-manager/internal/storage"
	"github.com/google/uuid"
	"time"
)

func CreatePendingTransaction() *storage.Transaction {
	t := storage.Transaction{
		Number:    uuid.New().String(),
		AppleId:   sql.NullString{String: "", Valid: false},
		Status:    storage.TransactionPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &t
}

func CreateEnrollOrder(customerId string, contractDate time.Time, number string) *storage.Order {
	o := storage.Order{
		CustomerId:   customerId,
		Number:       number,
		OrderType:    storage.OrderTypeEnroll,
		ContractDate: contractDate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return &o
}

func CreateDelivery(deliveryNumber string) *storage.Delivery {
	d := storage.Delivery{
		Number:    deliveryNumber,
		Status:    storage.DeliveryPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &d
}

func CreateDevices(devices []model.Device) []*storage.Device {
	var d []*storage.Device
	for _, device := range devices {
		d = append(d, &storage.Device{
			Imei:      device.Imei,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	return d
}
