package query

const (
	Asc  Order = "ASC"
	Desc Order = "DESC"
)

type (
	Order string

	Pagination struct {
		Page      uint
		Limit     uint
		SortBy    string
		SortOrder Order
	}
)

func (p Pagination) Offset() uint {
	return (p.Page - 1) * p.Limit
}
