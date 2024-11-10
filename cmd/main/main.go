package main

import (
	"github.com/illusory-server/accounts/internal/infra/config"
	"github.com/illusory-server/accounts/pkg/app"
)

func main() {
	a := app.Init("accounts", "0.0.1", "Accounts service").WithConfig(&config.Config{})

	err := a.Start()
	if err != nil {
		panic(err)
	}
}
