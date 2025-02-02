package main

import (
	"github.com/illusory-server/accounts/cmd/dependency"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/illusory-server/accounts/pkg/ecosystem"
)

func main() {
	dependencyFactory := dependency.NewFactory()

	app := ayaka.NewApp(&ayaka.Options{
		Name:              "Accounts",
		Description:       "Core accounts service",
		Version:           "0.0.1",
		ConfigInterceptor: ecosystem.AdapterParseConfigFromEnv,
		Container:         dependencyFactory.Container(),
	})

	err := ecosystem.StartWithCli(app, ecosystem.DefaultPrinter{})
	if err != nil {
		panic(err)
	}
}
