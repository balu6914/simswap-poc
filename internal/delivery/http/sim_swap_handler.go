package http

import (
	"encoding/json"
	"net/http"
	"simswap-poc/internal/domain"
	"simswap-poc/internal/usecase"

	"github.com/labstack/echo/v4"
)

type SimSwapHandler struct {
	usecase usecase.SimSwapUsecase
}

func NewSimSwapHandler(e *echo.Echo, uc usecase.SimSwapUsecase) {
	handler := &SimSwapHandler{usecase: uc}
	e.GET("/health", handler.HealthCheck)
	e.GET("/", handler.HealthCheck)
	e.POST("/sim-swap/v2/retrieve-date", handler.RetrieveSimSwapDate)
	e.POST("/sim-swap/v2/check", handler.CheckSimSwap)
}

func (h *SimSwapHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}

func (h *SimSwapHandler) RetrieveSimSwapDate(c echo.Context) error {
	var req domain.CreateSimSwapDate
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Generic400) // Invalid JSON or malformed request
	}

	result, errResp := h.usecase.RetrieveSimSwapDate(req)
	if errResp != nil {
		return c.JSON(errResp.Code, errResp) // Return the error response as-is
	}

	return c.JSON(http.StatusOK, result)
}

func (h *SimSwapHandler) CheckSimSwap(c echo.Context) error {
	var req domain.CreateCheckSimSwap
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Generic400) // Invalid JSON or malformed request
	}

	result, errResp := h.usecase.CheckSimSwap(req)
	if errResp != nil {
		return c.JSON(errResp.Code, errResp) // Return the error response as-is
	}

	return c.JSON(http.StatusOK, result)
}
