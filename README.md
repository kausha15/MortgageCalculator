# Quoter Technical Assessment

## Mortgage Calculator UI

Using a Frontend framework of your choice, build a BC mortgage calculator UI that makes use of the provided Go API Backend. The Backend API handles the business logic for calculating a BC mortgage, and considers CMHC insurance. Additional information about the calculation and restrictions can be found here:
https://www.ratehub.ca/cmhc-insurance-british-columbia.

Note: The Backend API may not have the same decimal precision as other mortgage calculators. Don't worry if the calculation is a few cents off of what other mortgage calculators show.

## Requirements

- The UI should include inputs for:
  - property price
  - down payment
  - annual interest rate
  - amortization period (5 year increments between 5 and 30 years)
  - payment schedule (accelerated bi-weekly, bi-weekly, monthly)
- A button to submit the input fields to the API
- After receiving a response from the API:
  - if the API response is successful, display payment per payment schedule
  - if the API response returns an error, display the error message

Bonus Points:

- Client-side input validation
- Unit tests

## Running the Go API Backend

Note: You can specify which port the API uses by changing `HTTP_PORT` if you already have something running on 8080.

```
$ cd calculator_api
$ go get ./...
$ HTTP_PORT=8080 go run main.go
```

Now, the API endpoint should be accessible at `http://localhost:8080/calculate`

## API Examples

#### Request

`POST /calculate`

#### Example Request

```
{
    "price": 80972,
    "down_payment": 4049,
    "annual_interest_rate": 5,
    "amortization_period": 20,
    "payment_schedule": "monthly"
}
```

#### Example Response

```
{
    "payment":527.96,
    "payment_schedule":"monthly"
}
```

Note: Please see POSTMAN collection `Quoter Assignment.postman_collection.josn` for further details

## Assignment Notes

Please Include any notes about your assignment here, including instructions for how to run your project.
