package storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
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
	Id        int64             `db:"id"`
	Number    string            `db:"number"`
	AppleId   sql.NullString    `db:"apple_id"`
	Status    TransactionStatus `db:"status"`
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt time.Time         `db:"updated_at"`
}

func SaveTransaction(db *sqlx.DB, tran *Transaction) (*Transaction, error) {
	stmt, err := db.PrepareNamed(
		`INSERT INTO transactions (number, apple_id, status, created_at, updated_at)
				VALUES (:number, :apple_id, :status, :created_at, :updated_at) RETURNING id`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&tran.Id, tran)
	if err != nil {
		return nil, err
	}

	return tran, nil
}

func SaveTransactionTx(db *sqlx.Tx, tran *Transaction) (*Transaction, error) {
	stmt, err := db.PrepareNamed(
		`INSERT INTO transactions (number, apple_id, status, created_at, updated_at)
				VALUES (:number, :apple_id, :status, :created_at, :updated_at) RETURNING id`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&tran.Id, tran)
	if err != nil {
		return nil, err
	}

	return tran, nil
}
