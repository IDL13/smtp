package main

import (
	"log"

	"github.com/IDL13/smtp/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
