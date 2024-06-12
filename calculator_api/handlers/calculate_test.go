package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"quoter_assignment/calculator_api/data"
	"testing"
)

func TestCalculate(t *testing.T) {
	t.Run("calculate integration test", func(t *testing.T) {
		t.Run("can calculate payment", func(t *testing.T) {
			mr := data.MortgageRequest{
				Price:        100_000,
				DownPayment:  5_000,
				InterestRate: 5.0,
				Amortization: 25,
				Schedule:     "monthly",
			}

			reqJson, err := json.Marshal(mr)
			if err != nil {
				t.Error("should not have failed to serialize")
			}

			reader := bytes.NewReader(reqJson)

			l := log.New(os.Stdout, "calculate-api ", log.LstdFlags)
			calculator := NewCalculator(l)
			req := httptest.NewRequest(http.MethodPost, "/calculate", reader)
			res := httptest.NewRecorder()

			calculator.Calculate(res, req)

			assert.Equal(t, res.Code, 200)
			assert.Equal(t, res.Body.String(), "{\"payment\":577.57,\"payment_schedule\":\"monthly\"}\n")
		})

		t.Run("400 if missing required field", func(t *testing.T) {
			mr := data.MortgageRequest{
				Price:        100_000,
				DownPayment:  5_000,
				Amortization: 25,
				Schedule:     "monthly",
			}

			reqJson, err := json.Marshal(mr)
			if err != nil {
				t.Error("should not have failed to serialize")
			}

			reader := bytes.NewReader(reqJson)

			l := log.New(os.Stdout, "calculate-api ", log.LstdFlags)
			calculator := NewCalculator(l)
			req := httptest.NewRequest(http.MethodPost, "/calculate", reader)
			res := httptest.NewRecorder()

			calculator.Calculate(res, req)

			assert.Equal(t, res.Code, 400)
		})

		t.Run("400 if validation fails", func(t *testing.T) {
			mr := data.MortgageRequest{
				Price:        100_000,
				DownPayment:  5_000,
				InterestRate: 5.0,
				Amortization: 25,
				Schedule:     "",
			}

			reqJson, err := json.Marshal(mr)
			if err != nil {
				t.Error("should not have failed to serialize")
			}

			reader := bytes.NewReader(reqJson)

			l := log.New(os.Stdout, "calculate-api ", log.LstdFlags)
			calculator := NewCalculator(l)
			req := httptest.NewRequest(http.MethodPost, "/calculate", reader)
			res := httptest.NewRecorder()

			calculator.Calculate(res, req)

			assert.Equal(t, res.Code, 400)
		})

		t.Run("400 if no request body", func(t *testing.T) {
			l := log.New(os.Stdout, "calculate-api ", log.LstdFlags)
			calculator := NewCalculator(l)
			req := httptest.NewRequest(http.MethodPost, "/calculate", nil)
			res := httptest.NewRecorder()

			calculator.Calculate(res, req)

			assert.Equal(t, res.Code, 400)
			assert.Equal(t, res.Body.String(), "invalid request body\n")
		})

		t.Run("400 if no invalid request body", func(t *testing.T) {
			anon := struct {
				name     string
				password string
			}{
				name:     "test",
				password: "pass",
			}

			reqJson, err := json.Marshal(anon)
			if err != nil {
				t.Error("should not have failed to serialize")
			}

			reader := bytes.NewReader(reqJson)
			l := log.New(os.Stdout, "calculate-api ", log.LstdFlags)
			calculator := NewCalculator(l)
			req := httptest.NewRequest(http.MethodPost, "/calculate", reader)
			res := httptest.NewRecorder()

			calculator.Calculate(res, req)

			assert.Equal(t, res.Code, 400)
		})
	})

}
