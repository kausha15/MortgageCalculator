package services

import (
	"github.com/stretchr/testify/assert"
	"quoter_assignment/calculator_api/data"
	"testing"
)

func TestMortgage(t *testing.T) {
	t.Run("get CHMH rate", func(t *testing.T) {
		t.Run("loan-to-value of 65%", func(t *testing.T) {
			req := newSimpleMortgageRequest()
			req.Price = 100
			req.DownPayment = 35

			assert.Equal(t, 0.006, getCMHCRate(req))
		})

		t.Run("loan-to-value of 75%", func(t *testing.T) {
			req := newSimpleMortgageRequest()
			req.Price = 100
			req.DownPayment = 25

			assert.Equal(t, 0.017, getCMHCRate(req))
		})

		t.Run("loan-to-value of 80%", func(t *testing.T) {
			req := newSimpleMortgageRequest()
			req.Price = 100
			req.DownPayment = 20

			assert.Equal(t, 0.024, getCMHCRate(req))
		})

		t.Run("loan-to-value of 85%", func(t *testing.T) {
			req := newSimpleMortgageRequest()
			req.Price = 100
			req.DownPayment = 15

			assert.Equal(t, 0.028, getCMHCRate(req))
		})

		t.Run("loan-to-value of 90%", func(t *testing.T) {
			req := newSimpleMortgageRequest()
			req.Price = 100
			req.DownPayment = 10

			assert.Equal(t, 0.031, getCMHCRate(req))
		})

		t.Run("loan-to-value of 95%", func(t *testing.T) {
			req := newSimpleMortgageRequest()
			req.Price = 100
			req.DownPayment = 5

			assert.Equal(t, 0.04, getCMHCRate(req))
		})
	})

	t.Run("validate and return CMHC", func(t *testing.T) {
		t.Run("under 500k", func(t *testing.T) {
			t.Run("under 20% down", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 499_999
				req.DownPayment = 25_000

				cmhc, err := validateAndReturnCMHC(req)
				if err != nil {
					t.Error("should not have errored")
				}
				if cmhc == 0.0 {
					t.Error("cmhc should not be 0.0")
				}
			})

			t.Run("20% down", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 499_999
				req.DownPayment = 100_000

				cmhc, err := validateAndReturnCMHC(req)
				if err != nil {
					t.Error("should not have errored")
				}
				if cmhc != 0.0 {
					t.Error("cmhc should be 0.0")
				}
			})

			t.Run("under 5% down", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 499_999
				req.DownPayment = 20_000

				_, err := validateAndReturnCMHC(req)
				if err == nil {
					t.Error("should have errored")
				}
			})

			t.Run("under 20% down and over 25 years", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 499_999
				req.DownPayment = 25_000
				req.Amortization = 30

				_, err := validateAndReturnCMHC(req)
				if err == nil {
					t.Error("should have errored")
				}
			})
		})

		t.Run("over 500k", func(t *testing.T) {
			t.Run("over 7.5% under 20% down", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 999_999
				req.DownPayment = 75_000

				cmhc, err := validateAndReturnCMHC(req)
				if err != nil {
					t.Error("should not have errored")
				}
				if cmhc == 0.0 {
					t.Error("cmhc should not be 0.0")
				}
			})

			t.Run("20% down", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 999_999
				req.DownPayment = 200_000

				cmhc, err := validateAndReturnCMHC(req)
				if err != nil {
					t.Error("should not have errored")
				}
				if cmhc != 0.0 {
					t.Error("cmhc should be 0.0")
				}
			})

			t.Run("under 20% down and over 25 years", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 999_999
				req.DownPayment = 75_000
				req.Amortization = 30

				_, err := validateAndReturnCMHC(req)
				if err == nil {
					t.Error("should have errored")
				}
			})

			t.Run("under 7.5% down", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 999_999
				req.DownPayment = 50_000

				_, err := validateAndReturnCMHC(req)
				if err == nil {
					t.Error("should have errored")
				}
			})

		})

		t.Run("1 million", func(t *testing.T) {
			t.Run("20% down", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 1_000_000
				req.DownPayment = 200_000

				cmhc, err := validateAndReturnCMHC(req)
				if err != nil {
					t.Error("should not have errored")
				}
				if cmhc != 0.0 {
					t.Error("cmhc should be 0.0")
				}
			})

			t.Run("under 20% down", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 1_000_000
				req.DownPayment = 100_000

				_, err := validateAndReturnCMHC(req)
				if err == nil {
					t.Error("should have errored")
				}
			})
		})
	})

	t.Run("Calculate", func(t *testing.T) {
		t.Run("under 500k", func(t *testing.T) {
			t.Run("under 20% down monthly", func(t *testing.T) {
				req := newSimpleMortgageRequest()

				res, err := Calculate(req)
				if err != nil {
					t.Error("should not have got an error")
				}

				assert.Equal(t, 577.57, res)
			})

			t.Run("under 20% down bi-weekly", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Schedule = "bi-weekly"

				res, err := Calculate(req)
				if err != nil {
					t.Error("should not have got an error")
				}

				assert.Equal(t, 266.57, res)
			})

			t.Run("under 20% down accel-bi-weekly", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Schedule = "accel-bi-weekly"

				res, err := Calculate(req)
				if err != nil {
					t.Error("should not have got an error")
				}

				assert.Equal(t, 288.79, res)
			})

			t.Run("20% down monthly", func(t *testing.T) {
				req := newSimpleMortgageRequest()
				req.Price = 499_999
				req.DownPayment = 100_000

				res, err := Calculate(req)
				if err != nil {
					t.Error("should not have got an error")
				}

				assert.Equal(t, 2338.35, res)
			})

		})

		t.Run("over 500k under 1 million with 20% monthly", func(t *testing.T) {
			req := newSimpleMortgageRequest()
			req.Price = 999_999
			req.DownPayment = 200_000

			res, err := Calculate(req)
			if err != nil {
				t.Error("should not have got an error")
			}

			assert.Equal(t, 4676.71, res)
		})

		t.Run("1 million with 20% monthly", func(t *testing.T) {
			req := newSimpleMortgageRequest()
			req.Price = 1_200_000
			req.DownPayment = 240_000

			res, err := Calculate(req)
			if err != nil {
				t.Error("should not have got an error")
			}

			assert.Equal(t, 5612.06, res)
		})
	})
}

func newSimpleMortgageRequest() data.MortgageRequest {
	return data.MortgageRequest{
		Price:        100_000,
		DownPayment:  5_000,
		InterestRate: 5.0,
		Amortization: 25,
		Schedule:     "monthly",
	}
}
