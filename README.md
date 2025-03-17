# SIM Swap POC

This is a Proof of Concept (POC) for a SIM Swap API built using Go and HarperDB.

## Project Structure

    .
        ├── config
        │ └── config.go
        ├── config.yaml
        ├── database
        │ ├── database.go
        │ └── harperdb.go
        ├── entities
        │ └── simswap.go
        ├── go.mod
        ├── go.sum
        ├── handlers
        │ ├── simswapHandler.go
        │ ├── simswapHttp.go
        │ └── simswapResponse.go
        ├── main.go
        ├── models
        │ └── simswap.go
        ├── repositories
        │ ├── simswapHarperDBRepository.go
        │ └── simswapRepository.go
        ├── server
        │ ├── echoServer.go
        │ └── server.go
        ├── tests
        │ └── simswap_test.go
        └── usecases
        ├── simswapUsecase.go
        └── simswapUsecaseImpl.go


## How to Run

1. Clone the repository:
   
   git clone https://github.com/balu6914/simswap-poc.git


2. Run the application:

    go run main.go

## API Endpoints
Retrieve SIM Swap Date: POST /retrieve-date

Check SIM Swap in Period: POST /check

## API Endpoints

Retrieve SIM Swap Date: POST /retrieve-date

Check SIM Swap in Period: POST /check

