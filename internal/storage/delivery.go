package storage

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type DeliveryStatus string

const (
	DeliveryPending    DeliveryStatus = "pending"
	DeliveryProcessing DeliveryStatus = "processing"
	DeliveryCompleted  DeliveryStatus = "completed"
	DeliveryFailed     DeliveryStatus = "failed"
)

type Delivery struct {
	Id        int            `db:"id"`
	Number    string         `db:"number"`
	OrderId   int64          `db:"order_id"`
	Status    DeliveryStatus `db:"status"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

func SaveDelivery(db *sqlx.DB, delivery *Delivery) (*Delivery, error) {
	stmt, err := db.PrepareNamed(
		`INSERT INTO deliveries (number, order_id, status, created_at, updated_at)
				VALUES (:number, :order_id, :status, :created_at, :updated_at) RETURNING id`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&delivery.Id, delivery)
	if err != nil {
		return nil, err
	}

	return delivery, nil
}

func SaveDeliveryTx(db *sqlx.Tx, delivery *Delivery) (*Delivery, error) {
	stmt, err := db.PrepareNamed(
		`INSERT INTO deliveries (number, order_id, status, created_at, updated_at)
				VALUES (:number, :order_id, :status, :created_at, :updated_at) RETURNING id`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&delivery.Id, delivery)
	if err != nil {
		return nil, err
	}

	return delivery, nil
}
