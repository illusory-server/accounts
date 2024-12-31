package grpc_job

import (
	"context"
	"encoding/json"
	"github.com/illusory-server/accounts/pkg/ecosystem"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

type testOut struct {
	out []byte
}

func (t *testOut) Write(p []byte) (n int, err error) {
	t.out = p
	return len(p), nil
}

func TestAppLoggerFromSlog(t *testing.T) {
	out := &testOut{}
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	log := ecosystem.NewAppLoggerWithSlog(logger)

	log.Debug(ctx, "kekw", map[string]interface{}{
		"foo": "bar",
		"baz": "qux",
	})

	var data map[string]interface{}
	err := json.Unmarshal(out.out, &data)
	assert.NoError(t, err)

	assert.Equal(t, "bar", data["foo"])
	assert.Equal(t, "qux", data["baz"])
	assert.Equal(t, "kekw", data["msg"])
	assert.Equal(t, "DEBUG", data["level"])

	log.Info(ctx, "kekw", map[string]interface{}{
		"foo": "bar",
		"baz": "qux",
		"hoho": map[string]interface{}{
			"hehe": "lol",
			"kek":  "xd",
		},
	})
	data = map[string]interface{}{}
	err = json.Unmarshal(out.out, &data)
	assert.NoError(t, err)

	assert.Equal(t, "bar", data["foo"])
	assert.Equal(t, "qux", data["baz"])
	assert.Equal(t, "kekw", data["msg"])
	assert.Equal(t, map[string]interface{}{"hehe": "lol", "kek": "xd"}, data["hoho"])
	assert.Equal(t, "INFO", data["level"])

	log.Warn(ctx, "kekw", map[string]interface{}{
		"foo": "bar",
		"baz": "qux",
		"hoho": map[string]interface{}{
			"hehe": "lol",
			"kek":  "xd",
		},
	})
	data = map[string]interface{}{}
	err = json.Unmarshal(out.out, &data)
	assert.NoError(t, err)

	assert.Equal(t, "bar", data["foo"])
	assert.Equal(t, "qux", data["baz"])
	assert.Equal(t, "kekw", data["msg"])
	assert.Equal(t, map[string]interface{}{"hehe": "lol", "kek": "xd"}, data["hoho"])
	assert.Equal(t, "WARN", data["level"])

	log.Error(ctx, "kekw", map[string]interface{}{
		"foo": "bar",
		"baz": "qux",
		"hoho": map[string]interface{}{
			"hehe": "lol",
			"kek":  "xd",
		},
	})
	data = map[string]interface{}{}
	err = json.Unmarshal(out.out, &data)
	assert.NoError(t, err)

	assert.Equal(t, "bar", data["foo"])
	assert.Equal(t, "qux", data["baz"])
	assert.Equal(t, "kekw", data["msg"])
	assert.Equal(t, map[string]interface{}{"hehe": "lol", "kek": "xd"}, data["hoho"])
	assert.Equal(t, "ERROR", data["level"])

	log.Info(ctx, "kekw2", nil)
	data = map[string]interface{}{}
	err = json.Unmarshal(out.out, &data)
	assert.NoError(t, err)

	assert.Equal(t, nil, data["foo"])
	assert.Equal(t, nil, data["baz"])
	assert.Equal(t, "kekw2", data["msg"])
	assert.Equal(t, nil, data["hoho"])
	assert.Equal(t, "INFO", data["level"])
}
