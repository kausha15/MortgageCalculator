package services

import (
	"errors"
	"math"
	"quoter_assignment/calculator_api/data"
)

// Calculate calculates the payment using the below formula
// M = P * r(1+r)^n) / (1+r)^n - 1
// M = payment per payment schedule
// P = principal
// r = per payment schedule interest rate
// n = total number of payments over the amortization period
func Calculate(mr data.MortgageRequest) (float64, error) {
	cmhc, err := validateAndReturnCMHC(mr)
	if err != nil {
		return 0, err
	}

	r := (mr.InterestRate * 0.01) / 12 // 12 months
	n := float64(mr.Amortization) * 12
	P := (mr.Price - mr.DownPayment) * (1 + cmhc)

	monthly := P * ((r * math.Pow(1+r, n)) / (math.Pow(1+r, n) - 1))
	M := monthly

	if mr.Schedule == "bi-weekly" {
		M = monthly * 12 / 26
	} else if mr.Schedule == "accel-bi-weekly" {
		M = monthly / 2
	}

	return math.Round(M*100) / 100, nil
}

// validateAndReturnCMHC validates whether the specific rules below have been met and returns CMHC.
// 1. down payment for $500k and under must be at least 5%
// 2. down payment for over $500k and under $1 million must be 5% on the first $500k and 10% on the rest
// 3. down payment for $1 million and up must be at least 20%
// 4. under $1 million and under 20% will require CMHC
// 5. largest amortization if CMHC is required is 25 years
func validateAndReturnCMHC(mr data.MortgageRequest) (float64, error) {
	var requireCMHC bool
	var cmhc float64
	downPaymentPerc := mr.DownPayment / mr.Price

	if mr.Price <= 500_000 {
		if downPaymentPerc < 0.05 {
			return 0.0, errors.New("down payment must be at least 5%")
		}

		if downPaymentPerc < 0.2 {
			requireCMHC = true
		}
	} else if mr.Price < 1_000_000 {
		uh := mr.Price - 500_000
		minDownPaymentPerc := (500_000 * 0.05) + (uh * 0.1)
		if mr.DownPayment < minDownPaymentPerc {
			return 0.0, errors.New("down payment for over $500,000 and under $1 million must be 5% on the first $500,000 and 10% on the rest")
		}

		if downPaymentPerc < 0.2 {
			requireCMHC = true
		}
	} else {
		// 1 million or above
		if downPaymentPerc < 0.2 {
			return 0.0, errors.New("down payment for $1 million or more must be 20%")
		}
	}

	// if CMHC is required calculate it
	if requireCMHC {
		if mr.Amortization > 25 {
			return 0.0, errors.New("maximum amortization period with less than 20% down is 25 years")
		}
		cmhc = getCMHCRate(mr)
	}

	return cmhc, nil
}

// calculateCMHC calculates the mortgage default insurance rate also known as CMHC
func getCMHCRate(mr data.MortgageRequest) float64 {
	// loan-to-value - ratio of a loan to the value of the asset
	ltv := (mr.Price - mr.DownPayment) / mr.Price
	var percentRate float64

	if ltv <= 0.65 {
		percentRate = 0.006
	} else if ltv <= 0.75 {
		percentRate = 0.017
	} else if ltv <= 0.80 {
		percentRate = 0.024
	} else if ltv <= 0.85 {
		percentRate = 0.028
	} else if ltv <= 0.9 {
		percentRate = 0.031
	} else {
		percentRate = 0.04
	}

	return percentRate
}
