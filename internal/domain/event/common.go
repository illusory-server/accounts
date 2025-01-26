package event

import "time"

type (
	Type string

	Event interface {
		Type() Type
		Timestamp() time.Time
	}
)
