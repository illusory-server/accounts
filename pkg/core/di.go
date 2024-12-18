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

type SyncContainer struct {
	container *dig.Container
	mu        *sync.Mutex
}

func NewSyncContainer(container *dig.Container) *SyncContainer {
	return &SyncContainer{
		container: container,
		mu:        &sync.Mutex{},
	}
}

func (s *SyncContainer) Invoke(function interface{}, opts ...dig.InvokeOption) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.Invoke(function, opts...)
}

func (s *SyncContainer) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.String()
}

func (s *SyncContainer) Scope(name string, opts ...dig.ScopeOption) *dig.Scope {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.Scope(name, opts...)
}

func (s *SyncContainer) Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.Provide(constructor, opts...)
}

func (s *SyncContainer) Decorate(decorator interface{}, opts ...dig.DecorateOption) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.container.Decorate(decorator, opts...)
}
