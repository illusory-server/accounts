package ayaka

import (
	"context"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

type testLogger struct {
	levels   []string
	messages []string
	infos    []map[string]any
	mutex    *sync.Mutex
}

func (t *testLogger) Debug(_ context.Context, message string, info map[string]any) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.levels = append(t.levels, "debug")
	t.messages = append(t.messages, message)
	t.infos = append(t.infos, info)
}

func (t *testLogger) Info(_ context.Context, message string, info map[string]any) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.levels = append(t.levels, "info")
	t.messages = append(t.messages, message)
	t.infos = append(t.infos, info)
}

func (t *testLogger) Warn(_ context.Context, message string, info map[string]any) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.levels = append(t.levels, "warn")
	t.messages = append(t.messages, message)
	t.infos = append(t.infos, info)
}

func (t *testLogger) Error(_ context.Context, message string, info map[string]any) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.levels = append(t.levels, "error")
	t.messages = append(t.messages, message)
	t.infos = append(t.infos, info)
}

func newTestLogger() *testLogger {
	return &testLogger{
		levels:   make([]string, 0, 10),
		messages: make([]string, 0, 10),
		infos:    make([]map[string]any, 0, 10),
		mutex:    &sync.Mutex{},
	}
}

func TestLogger(t *testing.T) {
	logger := newTestLogger()

	ayaka.NewApp(&ayaka.Options{
		Name:        "my-app",
		Description: "my-app description testing",
		Version:     "1.0.0",
		Logger:      logger,
	}).WithConfig(&ayaka.Config{
		StartTimeout:    time.Second * 2,
		GracefulTimeout: time.Second * 3,
		Info: map[string]any{
			"test": "kek",
		},
	})

	assert.Equal(t, []string{"info"}, logger.levels)
	assert.Equal(t, []string{"add new config"}, logger.messages)
	assert.Equal(t, []map[string]any{
		{
			"config": &ayaka.Config{
				StartTimeout:    time.Second * 2,
				GracefulTimeout: time.Second * 3,
				Info: map[string]any{
					"test": "kek",
				},
			},
		},
	}, logger.infos)
}
