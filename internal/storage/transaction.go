package storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

type TransactionStatus string

const (
	TransactionPending    TransactionStatus = "pending"
	TransactionProcessing TransactionStatus = "processing"
	TransactionCompleted  TransactionStatus = "completed"
	TransactionFailed     TransactionStatus = "failed"
)

type Transaction struct {
	Id                 int64             `db:"id"`
	UUID               string            `db:"uuid"`
	AppleTransactionId sql.NullString    `db:"apple_transaction_id"`
	Status             TransactionStatus `db:"status"`
	CreatedAt          time.Time         `db:"created_at"`
	UpdatedAt          time.Time         `db:"updated_at"`
	Order              *Order            `db:"-"`
}

func SaveEnrollmentTx(db *sqlx.Tx, tran *Transaction) error {
	tran, err := SaveTransactionTx(db, tran)
	if err != nil {
		return errors.Wrap(err, "failed to save transaction")
	}

	tran.Order.TransactionId = tran.Id
	tran.Order, err = SaveOrderTx(db, tran.Order)
	if err != nil {
		return errors.Wrap(err, "failed to save order")
	}

	for _, device := range tran.Order.Devices {
		device.OrderId = tran.Order.Id
	}

	_, err = SaveDevicesTx(db, tran.Order.Devices)
	if err != nil {
		return errors.Wrap(err, "failed to save devices")
	}

	return nil
}

func SaveTransactionTx(db *sqlx.Tx, tran *Transaction) (*Transaction, error) {
	stmt, err := db.PrepareNamed(
		`INSERT INTO transactions (uuid, apple_transaction_id, status, created_at, updated_at)
				VALUES (:uuid, :apple_transaction_id, :status, :created_at, :updated_at) RETURNING id`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&tran.Id, tran)
	if err != nil {
		return nil, err
	}

	return tran, nil
}

func GetTransactionById(db *sqlx.DB, id int64) (*Transaction, error) {
	var tran Transaction
	err := db.Get(&tran, "SELECT * FROM transactions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &tran, nil
}

func UpdateTransactionStatusAndAppleId(db *sqlx.DB, tran *Transaction) error {
	_, err := db.Exec(
		`UPDATE transactions SET status = $1, apple_transaction_id = $2, updated_at = $3 WHERE id = $4`,
		tran.Status, tran.AppleTransactionId, tran.UpdatedAt, tran.Id,
	)
	if err != nil {
		return err
	}

	return nil
}
