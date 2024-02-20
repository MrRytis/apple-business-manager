package factory

import (
	"database/sql"
	"github.com/MrRytis/apple-business-manager/internal/model"
	"github.com/MrRytis/apple-business-manager/internal/storage"
	"github.com/google/uuid"
	"time"
)

func CreatePendingTransaction(req model.EnrollRequest) *storage.Transaction {
	var devices []*storage.Device
	for _, device := range req.Devices {
		devices = append(devices, &storage.Device{
			Imei:      device.Imei,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	order := storage.Order{
		UUID:         uuid.New().String(),
		CustomerId:   req.CustomerId,
		OrderType:    storage.OrderTypeEnroll,
		ContractDate: req.ContractDate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Devices:      devices,
	}

	tran := storage.Transaction{
		UUID:               uuid.New().String(),
		AppleTransactionId: sql.NullString{String: "", Valid: false},
		Status:             storage.TransactionPending,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		Order:              &order,
	}

	return &tran
}
