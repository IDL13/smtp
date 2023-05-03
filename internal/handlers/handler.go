package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) StartHandler(c echo.Context) error {
	err := c.String(http.StatusOK, "[SERVER STARTED]")
	if err != nil {
		return err
	}
	return nil
}
