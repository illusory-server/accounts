package vo

import (
	"github.com/pkg/errors"
)

const (
	Asc  QueryOrder = "ASC"
	Desc QueryOrder = "DESC"
)

type (
	QueryOrder string

	Query struct {
		page      uint
		limit     uint
		sortBy    string
		sortOrder QueryOrder
	}
)

func (q QueryOrder) Validate() error {
	if q == Asc || q == Desc || q == "" {
		return nil
	}
	return errors.New("invalid query order value")
}

func (q Query) Page() uint {
	return q.page
}

func (q Query) Limit() uint {
	return q.limit
}

func (q Query) SortBy() string {
	return q.sortBy
}

func (q Query) SortOrder() QueryOrder {
	return q.sortOrder
}

func NewQuery(page, limit uint, sortBy string, sortOrder QueryOrder) (Query, error) {
	if err := sortOrder.Validate(); err != nil {
		return Query{}, err
	}
	return Query{
		page:      page,
		limit:     limit,
		sortBy:    sortBy,
		sortOrder: sortOrder,
	}, nil
}

func (q Query) PaginationOffset() uint {
	return (q.page - 1) * q.limit
}
