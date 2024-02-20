package storage

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Device struct {
	Id        int       `db:"id"`
	Imei      string    `db:"imei"`
	OrderId   int64     `db:"order_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func SaveDevicesTx(db *sqlx.Tx, devices []*Device) ([]*Device, error) {
	var d []Device
	for _, device := range devices {
		d = append(d, *device)
	}

	_, err := db.NamedExec(
		`INSERT INTO devices (imei, order_id, created_at, updated_at)
				VALUES (:imei, :delivery_id, :created_at, :updated_at)`, d)

	if err != nil {
		return nil, err
	}

	return devices, nil
}
