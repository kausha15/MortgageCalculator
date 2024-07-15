## Mortgage Calculator App

Mortgage Calculator App that makes use of the Go API Backend. The Backend API handles the business logic for calculating a BC mortgage, and considers CMHC insurance. Additional information about the calculation and restrictions can be found here:
https://www.ratehub.ca/cmhc-insurance-british-columbia.

This app also involves, 
 - API error handling
 - react form input validations 
 - unit tests using jest
 - performance review with lighthouse report


# Tech Stack : 
- Frontend : 
  - React
  - JavaScript
  - axios
- Backend : 
  - Go
- Unit Tests : 
  - Jest

# steps to run tests and application

## Running the Go API Backend

Note: You can specify which port the API uses by changing `HTTP_PORT` if you already have something running on 8080.

```
$ cd calculator_api
$ go get ./...
$ HTTP_PORT=8080 go run main.go
```

Now, the API endpoint should be accessible at `http://localhost:8080/calculate`

## Running the react UI

To run the application,

```
$ cd calculator_ui
$ npm install 
$ npm start

```
Now, the UI should be accessible at `http://localhost:3000`

## Running the Jest unit tests

```
$ cd calculator_ui
$ npm run test

```

## Overview of the solution

This Application is built for BC Mortgage Calculator using React, Javascript and Bootstrap.
As per the requirements, UI has a form with following fields with mentioned validations, 
 # Property Price 
    - This is requied field and number-only. 
    - step value is 1 and minimum value is 0.
    - Default value and placeholder is provided 
    - upon any change, percentage value is calculated and shown along with Down Payment
    - '$' sign is added as its a price value  
 # Down Payment
    - This is requied field and number-only.
    - step value is 1 and minimum value is 0. 
    - Default value and placeholder is provided  
    - upon any change, percentage value is calculated and shown along with Down Payment
    - percentage value is read-only
    - '$' sign is added as its a price value 
 # Annual Interest Rate
    - This is required field and accepts decimal value upto 2 digits (eg. 2.67 is correct and 2.6789 is not accepted)
    - validation for 2 decimal places happens on submit
    - '%' sign is added as its a rate 
    - accepts minimum value of 0
    - step value is 0.1
 # Amortization Period 
    - This is input field and read-only having default value of 20
    - acceptable values are between 5 to 30
    - values can only be updated with increment and decrement options on either side with a step value of 5
 # Schedule 
    - 3 schedules are presented 'Accelerated-Bi-Weekly | Bi-Weekly | Monthly' with radio buttons.
    - this is required field and 'Bi-Weekly' is selected by default



# Form Submission and Error Handling 
    - Form submission and trigger to API will happen only of all the fields have valid values.
    - Form values are sent to go api and ui is handled based on 3 scenarios,
    1. Successful response: 
       - response is parsed and 'payment_schedule' is conveted to pascal case and shown along with 'payment' value with green box.
    2. Unsuccessful resposne: 
       - In case of successful call execution and error in the resposne, error message is displayed with the red box.
    3. Call Failure:
       - In case of call failure, general error 'Something went wrong. Please try again in sometime.' is displyed with the red box. 


# Reason to calculate and display percentage besides down payment
Case where property price is '1200000' and down payment is '140000', api response is 'down payment for $1 million or more must be 20%'.
For better user experience, I have displayed the current percentage of down payment to property price and it will be updated real time when any of the two fields changes. 
By showing percentage value, it becomes easier for the user to estimate.

# Testing with Jest
Unit tests for basic operation is written such as,
1. is Calculator component rendering on the screen.
2. input fields are rendering on screen.
3. successful and unsuccessful form submit by mocking axios.

# Lighthouse Score for the application
Performance : 81%
Accessibility : 94%
Best Practices : 96%
SEO : 100%

Note : Screenshot of the Lighthouse report is available at the root directory




