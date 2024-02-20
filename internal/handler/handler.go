package handler

import (
	"github.com/MrRytis/apple-business-manager/internal/queue"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB     *sqlx.DB
	Rabbit *queue.Rabbit
}

func NewHandler(db *sqlx.DB, r *queue.Rabbit) *Handler {
	return &Handler{
		DB:     db,
		Rabbit: r,
	}
}
