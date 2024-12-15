package ayaka

import (
	"go.uber.org/dig"
	"sync"
)

type Container interface {
	Invoke(function interface{}, opts ...dig.InvokeOption) error
	String() string
	Scope(name string, opts ...dig.ScopeOption) *dig.Scope
	Provide(constructor interface{}, opts ...dig.ProvideOption) error
	Decorate(decorator interface{}, opts ...dig.DecorateOption) error
}

type syncContainer struct {
	container *dig.Container
	mu        *sync.Mutex
}

func newSyncContainer(container *dig.Container) *syncContainer {
	return &syncContainer{
		container: container,
		mu:        &sync.Mutex{},
	}
}

func (s *syncContainer) Invoke(function interface{}, opts ...dig.InvokeOption) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.Invoke(function, opts...)
}

func (s *syncContainer) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.String()
}

func (s *syncContainer) Scope(name string, opts ...dig.ScopeOption) *dig.Scope {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.Scope(name, opts...)
}

func (s *syncContainer) Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.Provide(constructor, opts...)
}

func (s *syncContainer) Decorate(decorator interface{}, opts ...dig.DecorateOption) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.Decorate(decorator, opts...)
}
