# SIM Swap Detection

> A service that integrates with a SIM Swap detection API to verify mobile numbers before processing sensitive user flows. Built using Go, Echo Framework, HarperDB, and Docker.


## Features

- SIM Swap status verification using external API
- Secure user onboarding with mobile validation
- Clean architecture (Go Clean Architecture)
- Lightweight HarperDB integration
- API testing with Postman
- CI/CD with GitHub Actions & Docker Hub


## Tech Stack

| Layer          | Technology                     |
|----------------|-------------------------------|
| Language       | Go (Golang)                    |
| Web Framework  | [Echo](https://echo.labstack.com/) |
| Database       | [HarperDB](https://harperdb.io/) |
| Deployment     | Docker                         |
| CI/CD          | GitHub Actions + Docker Hub    |
| API Testing    | Postman                        |


## Project Structure

├── Dockerfile
├── README.md
├── cmd
├── config
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── delivery
│   │   └── http
│   │       └── sim_swap_handler.go
│   ├── domain
│   │   ├── error_response.go
│   │   └── models.go
│   ├── repository
│   │   ├── harperdb.go
│   │   └── sim_swap_repo.go
│   └── usecase
│       └── sim_swap_usecase.go
├── main.go
└── simswap-poc

## Docker:
- Build & Run Locally
- docker build -t simswap-poc .
- docker run -p 9091:9091 --env-file .env simswap-poc
## CI/CD
This project uses GitHub Actions to:

- Run Go build & test on push

- Build & push Docker image to Docker Hub

## Docker image:
- balu1921/simswap-poc

