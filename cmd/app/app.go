package main

import (
	"github.com/IDL13/echo/pkg/utils"
	"github.com/IDL13/echo/smtp/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		utils.Loger(err)
	}

	a.Run()
}
