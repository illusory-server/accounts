package dependency

import (
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"go.uber.org/dig"
)

type Factory struct {
	di *dig.Container
}

func NewFactory() *Factory {
	di := ayaka.NewContainer(ayaka.NoopLogger{})
	return &Factory{
		di: di,
	}
}

func (f *Factory) Container() *dig.Container {
	return f.di
}
