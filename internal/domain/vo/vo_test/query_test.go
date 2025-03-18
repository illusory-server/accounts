package vo

import (
	"encoding/json"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codes"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery(t *testing.T) {
	t.Run("Should correct constructor", func(t *testing.T) {
		query, err := vo.NewQuery(2, 20, "name", vo.Asc)
		assert.NoError(t, err)
		assert.Equal(t, uint(2), query.Page())
		assert.Equal(t, uint(20), query.Limit())
		assert.Equal(t, "name", query.SortBy())
		assert.Equal(t, vo.Asc, query.SortOrder())

		query, err = vo.NewQuery(2, 20, "name", "")
		assert.NoError(t, err)
		assert.Equal(t, uint(2), query.Page())
		assert.Equal(t, uint(20), query.Limit())
		assert.Equal(t, "name", query.SortBy())
		assert.Equal(t, vo.QueryOrder(""), query.SortOrder())
	})

	t.Run("Should correct error with incorrect params", func(t *testing.T) {
		query, err := vo.NewQuery(5, 20, "", "incorrect order value")
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, errx.Code(err))
		assert.Equal(t, vo.Query{}, query)

		query, err = vo.NewQuery(5, 20, "", "k")
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, errx.Code(err))
		assert.Equal(t, vo.Query{}, query)
	})

	t.Run("Should correct pagination offset", func(t *testing.T) {
		query, err := vo.NewQuery(1, 20, "name", "")
		assert.NoError(t, err)
		assert.Equal(t, uint(0), query.PaginationOffset())

		query, err = vo.NewQuery(2, 20, "name", "")
		assert.NoError(t, err)
		assert.Equal(t, uint(20), query.PaginationOffset())

		query, err = vo.NewQuery(3, 20, "name", "")
		assert.NoError(t, err)
		assert.Equal(t, uint(40), query.PaginationOffset())

		query, err = vo.NewQuery(4, 20, "name", "")
		assert.NoError(t, err)
		assert.Equal(t, uint(60), query.PaginationOffset())
	})

	t.Run("Should correct marshal json", func(t *testing.T) {
		query, err := vo.NewQuery(2, 20, "name", vo.Asc)
		assert.NoError(t, err)
		js, err := json.Marshal(query)
		assert.NoError(t, err)
		assert.Equal(t, "{\"limit\":20,\"page\":2,\"sort_by\":\"name\",\"sort_order\":\"ASC\"}", string(js))
	})
}
