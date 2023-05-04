package main

import (
	"log"

	"smtp_server/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
