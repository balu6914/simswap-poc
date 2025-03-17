package server

import (
	"github.com/labstack/echo/v4"
	"simswap-poc/handlers"
	"simswap-poc/usecases"
	"simswap-poc/repositories"
	"simswap-poc/database"
	"simswap-poc/config"
)

func StartServer(cfg *config.Config) {
	e := echo.New()

	// Initialize HarperDB client
	harperDBClient := database.NewHarperDBClient(cfg.HarperDB.URL, cfg.HarperDB.Username, cfg.HarperDB.Password)

	// Initialize repository, use case, and handler
	repo := repositories.NewSimSwapHarperDBRepository(harperDBClient)
	usecase := usecases.NewSimSwapUsecase(repo)
	handler := handlers.NewSimSwapHandler(usecase)

	// Routes
	e.POST("/retrieve-date", handler.RetrieveSimSwapDate)
	e.POST("/check", handler.CheckSimSwap)

	// Start server
	e.Start(cfg.Server.Port)
}
