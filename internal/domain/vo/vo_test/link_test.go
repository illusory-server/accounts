package vo

import (
	"encoding/json"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVoLink(t *testing.T) {
	t.Run("Should correct constructor", func(t *testing.T) {
		l := "https://joska.com"
		link, err := vo.NewLink(l)
		assert.NoError(t, err)
		assert.Equal(t, l, link.Value())
	})

	t.Run("Should correct error with incorrect url", func(t *testing.T) {
		link, err := vo.NewLink("no url xdddd random text")
		assert.Error(t, err)
		assert.Equal(t, vo.Link{}, link)

		link, err = vo.NewLink("empty")
		assert.Error(t, err)
		assert.Equal(t, vo.Link{}, link)

	})

	t.Run("Should correct marshal json", func(t *testing.T) {
		l := "https://joska.com"
		link, err := vo.NewLink(l)
		assert.NoError(t, err)
		assert.Equal(t, l, link.Value())

		res, err := json.Marshal(link)
		assert.NoError(t, err)
		assert.Equal(t, "\""+l+"\"", string(res))
	})
}
