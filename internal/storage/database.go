package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func New() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTransaction(db *sqlx.DB) (*sqlx.Tx, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	return tx, nil
}
