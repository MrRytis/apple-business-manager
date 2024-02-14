package model

import (
	"github.com/MrRytis/apple-business-manager/internal/validator"
	"time"
)

type EnrollRequest struct {
	Number         string    `json:"number"`
	CustomerId     string    `json:"customer_id"`
	DeliveryNumber string    `json:"delivery_number"`
	ContractDate   time.Time `json:"contract_date"`
	ShipDate       time.Time `json:"ship_date"`
	Devices        []Device  `json:"items"`
}

type Device struct {
	Imei string `json:"imei"`
}

func (r *EnrollRequest) Validate() validator.ValidationErrors {
	errors := validator.ValidationErrors{}

	err := validator.ValidateString(r.Number, "number", true, 0, 32)
	if len(err) > 0 {
		errors.Errors = append(errors.Errors, err...)
	}

	err = validator.ValidateString(r.CustomerId, "customer_id", true, 0, 32)
	if len(err) > 0 {
		errors.Errors = append(errors.Errors, err...)
	}

	err = validator.ValidateString(r.DeliveryNumber, "delivery_number", true, 0, 32)
	if len(err) > 0 {
		errors.Errors = append(errors.Errors, err...)
	}

	err = validator.ValidateDateGreater(r.ShipDate, "ship_date", time.Now())
	if len(err) > 0 {
		errors.Errors = append(errors.Errors, err...)

	}

	for _, device := range r.Devices {
		err = validator.ValidateString(device.Imei, "imei", true, 4, 20)
		if len(err) > 0 {
			errors.Errors = append(errors.Errors, err...)
		}

	}

	return errors
}

type EnrollResponse struct {
	Number string `json:"number"`
}
