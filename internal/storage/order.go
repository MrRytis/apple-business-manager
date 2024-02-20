package storage

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type OrderType string

const (
	OrderTypeEnroll   OrderType = "OR"
	OrderTypeUnEnroll OrderType = "RE"
	OrderTypeOverride OrderType = "OV"
	OrderTypeVoid     OrderType = "VD"
)

type Order struct {
	Id            int64     `db:"id"`
	UUID          string    `db:"uuid"`
	CustomerId    string    `db:"customer_id"`
	OrderType     OrderType `db:"order_type"`
	ContractDate  time.Time `db:"contract_date"`
	TransactionId int64     `db:"transaction_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	Devices       []*Device `db:"-"`
}

func SaveOrderTx(db *sqlx.Tx, order *Order) (*Order, error) {
	stmt, err := db.PrepareNamed(
		`INSERT INTO orders (customer_id, uuid, order_type, contract_date, transaction_id, created_at, updated_at)
				VALUES (:customer_id, :uuid, :order_type, :contract_date, :transaction_id, :created_at, :updated_at) RETURNING id`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&order.Id, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
