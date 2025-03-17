package server

import (
	"simswap-poc/config"
	"simswap-poc/database"
	"simswap-poc/handlers"
	"simswap-poc/repositories"
	"simswap-poc/usecases"

	"github.com/labstack/echo/v4"
)

func StartServer(cfg *config.Config) {
	e := echo.New()

	// HarperDB client
	harperDBClient := database.NewHarperDBClient(cfg.HarperDB.URL, cfg.HarperDB.Username, cfg.HarperDB.Password)

	repo := repositories.NewSimSwapHarperDBRepository(harperDBClient) // Initialize repository, use case, and handler
	usecase := usecases.NewSimSwapUsecase(repo)
	handler := handlers.NewSimSwapHandler(usecase)

	// Routes
	e.POST("/retrieve-date", handler.RetrieveSimSwapDate)
	e.POST("/check", handler.CheckSimSwap)

	e.Start(cfg.Server.Port) //server start
}
