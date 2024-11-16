package main

import (
	"github.com/illusory-server/accounts/internal/infra/config"
	"github.com/illusory-server/accounts/pkg/app"
)

func main() {
	a := app.Init(&app.Options{Name: "Accounts", Description: "base accounts service", Version: "1.0.0"}).
		WithConfig(&config.Config{})

	err := a.Start()
	if err != nil {
		panic(err)
	}
}
