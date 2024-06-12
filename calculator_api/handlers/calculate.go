package handlers

import (
	"log"
	"net/http"
	"quoter_assignment/calculator_api/data"
	"quoter_assignment/calculator_api/services"
)

type Calculator struct {
	l *log.Logger
}

func NewCalculator(l *log.Logger) *Calculator {
	return &Calculator{l}
}

// Calculate handler for POST /calculate
func (c *Calculator) Calculate(rw http.ResponseWriter, r *http.Request) {
	var mr data.MortgageRequest

	err := mr.FromJSON(r.Body)
	if err != nil {
		c.l.Println("error deserializing calculator request", err)

		if err.Error() == "EOF" {
			http.Error(rw, "invalid request body", http.StatusBadRequest)
			return
		}

		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	c.l.Printf("calculate request: %#v", mr)

	err = mr.Validate()
	if err != nil {
		c.l.Println("error validating calculator request:", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := services.Calculate(mr)
	if err != nil {
		c.l.Println("error calculating result:", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	c.l.Printf("calculated result: %.2f", res)

	resp := data.MortgagePayment{
		Payment:  res,
		Schedule: mr.Schedule,
	}

	err = resp.ToJSON(rw)
	if err != nil {
		// we should never hit this but log just in case
		c.l.Println("error serializing result")
	}
}
