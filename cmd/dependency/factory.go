package dependency

import "go.uber.org/dig"

type Factory struct {
	di *dig.Container
}

func NewFactory() *Factory {
	di := dig.New()
	return &Factory{
		di: di,
	}
}

func (f *Factory) Container() *dig.Container {
	return f.di
}
