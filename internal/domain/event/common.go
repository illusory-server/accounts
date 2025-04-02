package event

import (
	"encoding/json"
	"time"
)

type (
	Type string

	Event interface {
		Type() Type
		Value() any
		Timestamp() time.Time
	}

	Events []Event
)

func (e Events) MarshalJSON() ([]byte, error) {
	result := make([]string, len(e))
	for i, event := range e {
		result[i] = string(event.Type())
	}
	return json.Marshal(result)
}
