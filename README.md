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

## Running the Go API Backend

Note: You can specify which port the API uses by changing `HTTP_PORT` if you already have something running on 8080.

```
$ cd calculator_api
$ go get ./...
$ HTTP_PORT=8080 go run main.go
```

Now, the API endpoint should be accessible at `http://localhost:8080/calculate`

## Running the Frontend

```
$ cd calculator_ui
$ npm install 
$ npm start
```

Now, the UI endpoint should be accessible at `http://localhost:3000`