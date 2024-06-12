package data

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
)

// MortgageRequest is the input required to calculate mortgage payment
type MortgageRequest struct {
	Price        float64 `json:"price" validate:"required,gt=0"`
	DownPayment  float64 `json:"down_payment" validate:"required,gt=0"`
	InterestRate float64 `json:"annual_interest_rate" validate:"required,gt=0,lte=100"`
	Amortization int     `json:"amortization_period" validate:"required,oneof=5 10 15 20 25 30"`
	Schedule     string  `json:"payment_schedule" validate:"required,oneof=accel-bi-weekly bi-weekly monthly"`
}

// FromJSON deserialize MortgageRequest
func (c *MortgageRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

// Validate implement validation
func (c *MortgageRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
