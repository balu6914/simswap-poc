package main

import (
	"log"
	"simswap-poc/internal/delivery/http"
	"simswap-poc/internal/repository"
	"simswap-poc/internal/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func main() {
	e := echo.New()

	repo := repository.NewSimSwapRepository()
	useCase := usecase.NewSimSwapUsecase(repo)

	http.NewSimSwapHandler(e, useCase)

	e.Logger.Fatal(e.Start(":9091"))
}
