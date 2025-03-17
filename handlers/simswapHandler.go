package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"simswap-poc/usecases"
)

type SimSwapHandler struct {
	Usecase usecases.SimSwapUsecase
}

func NewSimSwapHandler(usecase usecases.SimSwapUsecase) *SimSwapHandler {
	return &SimSwapHandler{Usecase: usecase}
}

func (h *SimSwapHandler) RetrieveSimSwapDate(c echo.Context) error {
	var request struct {
		PhoneNumber string `json:"phoneNumber"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	date, err := h.Usecase.RetrieveSimSwapDate(request.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"latestSimChange": date})
}

func (h *SimSwapHandler) CheckSimSwap(c echo.Context) error {
	var request struct {
		PhoneNumber string `json:"phoneNumber"`
		MaxAge      int    `json:"maxAge"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	swapped, err := h.Usecase.CheckSimSwap(request.PhoneNumber, request.MaxAge)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]bool{"swapped": swapped})
}
