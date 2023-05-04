package app

import (
	"fmt"
	"net"

	handler "smtp_server/internal/handlers"

	"smtp_server/pkg/api"
	smtp "smtp_server/pkg/smtp"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
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
	// fmt.Println("[SERVER STARTED]")
	// a.echo.Logger.Fatal(a.echo.Start(":8085"))

	s := grpc.NewServer()
	srv := &smtp.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	fmt.Println("[SERVER STARTED]")
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(l); err != nil {
		fmt.Println(err)
	}
}
