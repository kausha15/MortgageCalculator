package data

import (
	"encoding/json"
	"io"
)

// MortgagePayment response for the mortgage calculation
type MortgagePayment struct {
	Payment  float64 `json:"payment"`
	Schedule string  `json:"payment_schedule"`
}

// ToJSON serializes MortgagePayment
func (c *MortgagePayment) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
