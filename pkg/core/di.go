package ayaka

import (
	"go.uber.org/dig"
)

type Container interface {
	Invoke(function interface{}, opts ...dig.InvokeOption) error
	String() string
	Scope(name string, opts ...dig.ScopeOption) *dig.Scope
	Provide(constructor interface{}, opts ...dig.ProvideOption) error
	Decorate(decorator interface{}, opts ...dig.DecorateOption) error
}
