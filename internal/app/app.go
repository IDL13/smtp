package app

import (
	"fmt"

	handler "github.com/IDL13/smtp/internal/handlers"
	"github.com/labstack/echo/v4"
)

type App struct {
	h    *handler.Handler
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.h = handler.New()
	a.echo = echo.New()

	a.echo.GET("/", a.h.StartHandler)
	a.echo.POST("/smtp", a.h.SmtpHandler)

	return a, nil
}

func (a *App) Run() {
	fmt.Println("[SERVER STARTED]")
	a.echo.Logger.Fatal(a.echo.Start(":8080"))
}
