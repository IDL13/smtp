package handler

import (
	"log"
	"net/http"
	"net/smtp"

	"smtp_server/internal/config"
	"smtp_server/internal/unmarshal"

	"github.com/labstack/echo/v4"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct {
	u *unmarshal.Context
}

func (h *Handler) StartHandler(c echo.Context) error {
	err := c.String(http.StatusOK, "[SERVER STARTED]")
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) SmtpHandler(c echo.Context) error {
	h.u = unmarshal.New()

	cfg := config.GetConf()
	h.u.Unmarshal(c)

	auth := smtp.PlainAuth("", "Nabokov@gmail.com", cfg.GmailKey, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "Nabokov@gmail.com", []string{h.u.Email}, []byte(h.u.Msg))
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, "successful request")
}
