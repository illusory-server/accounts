package main

import (
	"fmt"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/illusory-server/accounts/pkg/ecosystem"
)

func main() {
	app := ayaka.NewApp(&ayaka.Options{
		Name:        "TestStartWithCli",
		Description: "TestStartWithCli case",
		Version:     "1.0.0",
		Container:   ayaka.NewContainer(ayaka.NoopLogger{}),
	})
	err := ecosystem.StartWithCli(app, nil)
	if err != nil {
		fmt.Println(err)
	}
}
