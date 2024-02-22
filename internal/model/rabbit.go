package model

import (
	"github.com/MrRytis/apple-business-manager/internal/validator"
)

type ConsumerResponse struct {
	QueueName string `json:"queue_name"`
	Count     int    `json:"count"`
}

type AddConsumerRequest struct {
	QueueName string `json:"queue_name"`
}

func (r *AddConsumerRequest) Validate() validator.ValidationErrors {
	errors := validator.ValidationErrors{}

	err := validator.ValidateString(r.QueueName, "queue_name", true, 1, 255)
	if len(err) > 0 {
		errors.Errors = append(errors.Errors, err...)
	}

	return errors
}
